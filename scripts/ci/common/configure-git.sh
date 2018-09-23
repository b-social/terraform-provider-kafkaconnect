#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

git config --global user.email "circleci@b-social.com"
git config --global user.name "Circle CI"
