#!/usr/bin/env sh

set -x
set -e
set -o pipefail

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"

${SCRIPT_DIR}/install-bash.sh
${SCRIPT_DIR}/install-build-tools.sh
${SCRIPT_DIR}/install-ruby.sh
