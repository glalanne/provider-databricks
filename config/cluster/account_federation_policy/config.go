package account_federation_policy

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
// TODO: waiting on protov6 support with Upjet https://github.com/crossplane/upjet/issues/372
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("databricks_account_federation_policy", func(r *config.Resource) {
		r.ShortGroup = "oauth"
		r.TerraformResource.Schema["oidc_policy"].Type = schema.TypeMap
		r.TerraformResource.Schema["oidc_policy"].Elem = &schema.Resource{
			Schema: map[string]*schema.Schema{
				"audiences": {
					Type:     schema.TypeList,
					Optional: true,
					Computed: false,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "The allowed token audiences, as specified in the 'aud' claim of federated tokens. The audience identifier is intended to represent the recipient of the token. Can be any non-empty string value. As long as the audience in the token matches at least one audience in the policy, the token is considered a match. If audiences is unspecified, defaults to your Databricks account id",
				},
				"issuer": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    false,
					Description: "The required token issuer, as specified in the 'iss' claim of federated tokens",
				},
				"jwks_json": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    false,
					Description: "The public keys used to validate the signature of federated tokens, in JWKS format. Most use cases should not need to specify this field. If jwks_uri and jwks_json are both unspecified (recommended), Databricks automatically fetches the public keys from your issuer’s well known endpoint. Databricks strongly recommends relying on your issuer’s well known endpoint for discovering public keys",
				},
				"jwks_uri": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    false,
					Description: "URL of the public keys used to validate the signature of federated tokens, in JWKS format. Most use cases should not need to specify this field. If jwks_uri and jwks_json are both unspecified (recommended), Databricks automatically fetches the public keys from your issuer’s well known endpoint. Databricks strongly recommends relying on your issuer’s well known endpoint for discovering public keys",
				},
				"subject": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    false,
					Description: "The required token subject, as specified in the subject claim of federated tokens. Must be specified for service principal federation policies. Must not be specified for account federation policies",
				},
				"subject_claim": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    false,
					Description: "The claim that contains the subject of the token. If unspecified, the default value is 'sub'",
				},
			},
		}
	})
}
