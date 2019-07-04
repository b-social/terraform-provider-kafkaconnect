#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

export TERM=xterm

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/../../.." && pwd )"
export GOPATH=$HOME/go

cd "$PROJECT_DIR"

git crypt unlock
git pull

./build release:prepare[minor]
./build version:bump[minor]

git push origin master --tags

./build release:perform
