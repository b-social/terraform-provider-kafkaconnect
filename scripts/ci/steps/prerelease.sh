#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

export TERM=xterm

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/../../.." && pwd )"

cd "$PROJECT_DIR"

git crypt unlock

./build version:bump[rc]

git push origin master --tags

fingerprint=$(gpg --with-colons --list-keys | awk -F: '/^pub/ { print $5 }')
export GPG_FINGERPRINT=$fingerprint

./build release:perform
