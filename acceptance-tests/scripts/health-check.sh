#!/bin/sh

set -o errexit
set -o nounset

if ! command curl >/dev/null 2>&1; then
    apk add curl
fi

curl --fail http://localhost:9000/healthy || exit 1
