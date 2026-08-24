#!/usr/bin/env bash
set -aeuo pipefail

echo "Executing pre-assert hook for Jobs..."

# get job id
JOB_ID=""
TIMEOUT_SECONDS=120
POLL_INTERVAL_SECONDS=5
ELAPSED_SECONDS=0
NAMESPACE="upbound-system"
JOB_NAME="serverless-hello-world"

while [[ -z "${JOB_ID}" || "${JOB_ID}" == "0" ]]; do
     if (( ELAPSED_SECONDS >= TIMEOUT_SECONDS )); then
          echo "Timed out waiting for a valid Job ID after ${TIMEOUT_SECONDS} seconds." >&2
          exit 1
     fi   

     JOB_ID=$(kubectl get jobs.compute.databricks.m.crossplane.io "${JOB_NAME}" \
          -n "${NAMESPACE}" -o jsonpath='{.status.atProvider.id}')

     if [[ -z "${JOB_ID}" || "${JOB_ID}" == "0" ]]; then
          sleep "${POLL_INTERVAL_SECONDS}"
          ((ELAPSED_SECONDS += POLL_INTERVAL_SECONDS))
     fi
done

echo "Job found with ID: ${JOB_ID}"


# We assume the UPTEST_CLOUD_CREDENTIALS is using a PAT
# Need to find a better way

DATABRICKS_HOST=$(echo $UPTEST_CLOUD_CREDENTIALS | jq -r .host)
DATABRICKS_TOKEN=$(echo $UPTEST_CLOUD_CREDENTIALS | jq -r .token)

curl --request POST "${DATABRICKS_HOST}/api/2.2/jobs/run-now" \
     --header "Authorization: Bearer ${DATABRICKS_TOKEN}" \
     --header "Content-Type: application/json" \
     --data "{ \"job_id\": \"${JOB_ID}\" }"

sleep 30