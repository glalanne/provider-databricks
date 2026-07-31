#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

# Keep git behavior deterministic and avoid dependence on user-global config.
export GIT_CONFIG_GLOBAL="${GIT_CONFIG_GLOBAL:-/dev/null}"
export GIT_CONFIG_NOSYSTEM="${GIT_CONFIG_NOSYSTEM:-1}"

BASE_REF="$(printf '%s' "${1:-main}" | tr -d '\r')"
HEAD_REF="$(printf '%s' "${2:-HEAD}" | tr -d '\r')"
CRD_DIR="${CRD_DIR:-package/crds}"

if [[ ! -f Makefile ]]; then
  echo "This script must be run from the repository root."
  exit 2
fi

python3 ./scripts/crd_breaking_check.py --base-ref "${BASE_REF}" --head-ref "${HEAD_REF}" --crd-dir "${CRD_DIR}"
