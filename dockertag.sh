#!/bin/bash

set -eu

readonly VERSION=$1

readonly GREEN=$'\x1b[32m'
readonly RED=$'\x1b[31m'
readonly RESET=$'\x1b[m'

now() {
  date "+%Y/%m/%d %H:%M:%S"
}

info() {
  echo -e "[$(now)] ${GREEN}INFO${RESET}: $*"
}

err() {
  echo -e "[$(now)]  ${RED}ERR${RESET}: $*"
}

if [[ ! "$VERSION" =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  err "tag version (VERSION = $VERSION) is illegal. tag format must be 'v0.0.0'."
  exit 1
fi

readonly IMAGE="jiro4989/build-deb-action:$VERSION"
info "IMAGE = $IMAGE"
docker build --no-cache -t "$IMAGE" --target runtime .
docker push "$IMAGE"
info "Successful docker build and push"
