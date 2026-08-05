#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating provider credential secret..."
${KUBECTL} -n crossplane-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: databricks.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF

echo "Creating a default cluster provider config..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: databricks.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: crossplane-system
      key: credentials
EOF

# Creat a new namespace for namespaced tests name upbound-system
echo "Creating a new namespace for namespaced tests..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: upbound-system
EOF