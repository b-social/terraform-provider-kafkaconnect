#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

apk --update add \
    ca-certificates \
    ruby=2.6.6 \
    ruby-bundler \
    ruby-json

echo 'gem: --no-document' > /etc/gemrc
