#!/usr/bin/env sh

set -x
set -e
set -o pipefail

SCRIPT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"

${SCRIPT_DIR}/install-bash.sh
${SCRIPT_DIR}/install-build-tools.sh
${SCRIPT_DIR}/install-ruby.sh
${SCRIPT_DIR}/install-dep.sh
${SCRIPT_DIR}/install-git-crypt.sh
${SCRIPT_DIR}/install-gpg-key.sh
${SCRIPT_DIR}/configure-git.sh
