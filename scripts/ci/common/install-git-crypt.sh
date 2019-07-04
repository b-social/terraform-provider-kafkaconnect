#!/usr/bin/env bash

[ -n "$DEBUG" ] && set -x
set -e
set -o pipefail

GIT_CRYPT_VERSION="0.6.0"
GIT_CRYPT_PACKAGE_URL="https://www.agwa.name/projects/git-crypt/downloads/git-crypt-${GIT_CRYPT_VERSION}.tar.gz"
GIT_CRYPT_SIGNATURE_URL="https://www.agwa.name/projects/git-crypt/downloads/git-crypt-${GIT_CRYPT_VERSION}.tar.gz.asc"

apk --update add \
    curl \
    git \
    gnupg \
    g++ \
    make \
    openssh \
    openssl \
    openssl-dev

for key in 0xEF5D84C1838F2EB6D8968C0410378EFC2080080C; do
    gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys "$key" ||
    gpg --keyserver pgp.mit.edu --recv-keys "$key" ||
    gpg --keyserver keyserver.pgp.com --recv-keys "$key" ||
    gpg --keyserver ha.pool.sks-keyservers.net --recv-keys "$key" ||
    gpg --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys "$key";
done

cd /var/tmp
curl -SLO ${GIT_CRYPT_PACKAGE_URL}
curl -SLO ${GIT_CRYPT_SIGNATURE_URL}
#gpg --batch --verify \
#    "git-crypt-${GIT_CRYPT_VERSION}.tar.gz.asc" \
#    "git-crypt-${GIT_CRYPT_VERSION}.tar.gz"

mkdir -p /usr/src/git-crypt
tar -xzf "git-crypt-${GIT_CRYPT_VERSION}.tar.gz" \
    -C /usr/src/git-crypt --strip-components=1

rm "git-crypt-${GIT_CRYPT_VERSION}.tar.gz.asc"
rm "git-crypt-${GIT_CRYPT_VERSION}.tar.gz"

cd /usr/src/git-crypt
make
make install PREFIX=/usr/local
