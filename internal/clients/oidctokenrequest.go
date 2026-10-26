package clients

import (
	"context"

	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/glalanne/provider-databricks/apis/namespaced/v1beta1"
)

const (
	// annotationAzureClientID and annotationAzureTenantID are the Azure
	// workload identity annotations on a service account naming the managed
	// identity it is federated with.
	annotationAzureClientID = "azure.workload.identity/client-id"
	annotationAzureTenantID = "azure.workload.identity/tenant-id"

	// authTypePAT makes the Databricks SDK send the token as a bearer token.
	// Azure Databricks accepts Entra ID access tokens this way.
	authTypePAT = "pat"

	errServiceAccountOtherNamespace = "serviceAccountRef.namespace %q is not allowed: a namespaced ProviderConfig can only use service accounts in its own namespace %q"
	errNoServiceAccountRef          = "credentials.serviceAccountRef is required when credentials.source is OIDCTokenRequest"
	errNoServiceAccountNamespace    = "credentials.serviceAccountRef.namespace is required on a cluster-scoped ProviderConfig"
	errNoHost                       = "host is required when credentials.source is OIDCTokenRequest"
	errUnsupportedTokenExchange     = "unsupported credentials.tokenExchange.type %q"
	errGetServiceAccount            = "cannot get service account %s/%s to read its workload identity annotations"
	errNoClientID                   = "no client ID: set credentials.tokenExchange.azureAD.clientID, or annotate service account %s/%s with " + annotationAzureClientID
	errNoTenantID                   = "no tenant ID: set credentials.tokenExchange.azureAD.tenantID, or annotate service account %s/%s with " + annotationAzureTenantID
	errServiceAccountToken          = "cannot obtain an access token for service account %s/%s"
)

// defaultTokenCache is shared by all reconciles in the process, so each
// identity is exchanged once per token lifetime rather than once per connect.
var defaultTokenCache = NewTokenCache()

// enrichLocalServiceAccountRef pins a namespaced ProviderConfig's
// serviceAccountRef to the managed resource's namespace (the ProviderConfig's
// own), rejecting any other. Like enrichLocalSecretRefs for secrets, this is
// what stops one namespace from using another namespace's identity.
func enrichLocalServiceAccountRef(pc *namespacedv1beta1.ProviderConfig, mg xpresource.Managed) error {
	ref := pc.Spec.Credentials.ServiceAccountRef
	if ref == nil {
		return nil
	}
	ns := mg.GetNamespace()
	if ref.Namespace != nil && *ref.Namespace != ns {
		return errors.Errorf(errServiceAccountOtherNamespace, *ref.Namespace, ns)
	}
	ref.Namespace = &ns
	return nil
}

// oidcTokenRequestAuth configures the Terraform provider with an access
// token obtained for the ProviderConfig's service account. The service
// account's namespace must already be resolved: by enrichLocalServiceAccountRef
// for a namespaced ProviderConfig, or set explicitly on a cluster-scoped one.
func oidcTokenRequestAuth(ctx context.Context, kube client.Client, pcSpec *namespacedv1beta1.ProviderConfigSpec, ps *terraform.Setup, cache *TokenCache) error {
	ref, err := serviceAccountRef(pcSpec)
	if err != nil {
		return err
	}
	ns, name := *ref.Namespace, ref.Name

	ex, err := newExchanger(ctx, kube, pcSpec, ns, name)
	if err != nil {
		return err
	}

	src := &ServiceAccountTokenSource{
		Client:    kube,
		Namespace: ns,
		Name:      name,
		Audiences: withAudience(ex.Audience(), ref.Audiences),
	}
	token, err := cache.Get(ctx, src, ex)
	if err != nil {
		return errors.Wrapf(err, errServiceAccountToken, ns, name)
	}

	ps.Configuration[keyHost] = *pcSpec.Host
	ps.Configuration[keyAuthType] = authTypePAT
	ps.Configuration[keyAuthToken] = token.Value
	if v := deref(pcSpec.AccountID); v != "" {
		ps.Configuration[keyAccountID] = v
	}
	return nil
}

// serviceAccountRef returns the validated service account reference. The CRDs'
// CEL rules enforce this too, but old API servers may not evaluate CEL, and a
// namespaced ProviderConfig's namespace is only filled in at resolve time.
func serviceAccountRef(pcSpec *namespacedv1beta1.ProviderConfigSpec) (*namespacedv1beta1.ServiceAccountReference, error) {
	ref := pcSpec.Credentials.ServiceAccountRef
	if ref == nil {
		return nil, errors.New(errNoServiceAccountRef)
	}
	if deref(ref.Namespace) == "" {
		return nil, errors.New(errNoServiceAccountNamespace)
	}
	if deref(pcSpec.Host) == "" {
		return nil, errors.New(errNoHost)
	}
	return ref, nil
}

// newExchanger returns the exchanger for the ProviderConfig's tokenExchange
// type, defaulting to AzureAD.
func newExchanger(ctx context.Context, kube client.Client, pcSpec *namespacedv1beta1.ProviderConfigSpec, ns, name string) (Exchanger, error) {
	exchangeType := namespacedv1beta1.TokenExchangeAzureAD
	if te := pcSpec.Credentials.TokenExchange; te != nil && te.Type != nil {
		exchangeType = *te.Type
	}
	// The exhaustive linter requires a case here for every TokenExchangeType;
	// default only catches a value this binary doesn't know.
	switch exchangeType {
	case namespacedv1beta1.TokenExchangeAzureAD:
		ex, err := azureADExchanger(ctx, kube, pcSpec, ns, name)
		if err != nil {
			return nil, err
		}
		return ex, nil
	default:
		return nil, errors.Errorf(errUnsupportedTokenExchange, exchangeType)
	}
}

// azureADExchanger builds the Entra ID exchanger. clientID and tenantID come
// from credentials.tokenExchange.azureAD, or default to the service account's
// workload identity annotations.
func azureADExchanger(ctx context.Context, kube client.Client, pcSpec *namespacedv1beta1.ProviderConfigSpec, ns, name string) (*AzureADExchanger, error) {
	var settings namespacedv1beta1.AzureADTokenExchange
	if te := pcSpec.Credentials.TokenExchange; te != nil && te.AzureAD != nil {
		settings = *te.AzureAD
	}
	clientID, tenantID := deref(settings.ClientID), deref(settings.TenantID)
	if clientID == "" || tenantID == "" {
		sa := &corev1.ServiceAccount{}
		if err := kube.Get(ctx, types.NamespacedName{Namespace: ns, Name: name}, sa); err != nil {
			return nil, errors.Wrapf(err, errGetServiceAccount, ns, name)
		}
		if clientID == "" {
			clientID = sa.Annotations[annotationAzureClientID]
		}
		if tenantID == "" {
			tenantID = sa.Annotations[annotationAzureTenantID]
		}
	}
	if clientID == "" {
		return nil, errors.Errorf(errNoClientID, ns, name)
	}
	if tenantID == "" {
		return nil, errors.Errorf(errNoTenantID, ns, name)
	}

	return &AzureADExchanger{ClientID: clientID, TenantID: tenantID, AuthorityHost: deref(settings.AuthorityHost)}, nil
}

// withAudience returns required followed by the extra audiences, without
// duplicates.
func withAudience(required string, extra []string) []string {
	out := []string{required}
	for _, a := range extra {
		if a != "" && a != required {
			out = append(out, a)
		}
	}
	return out
}

func deref[T ~string](s *T) string {
	if s == nil {
		return ""
	}
	return string(*s)
}
