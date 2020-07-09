#!/bin/sh

set -eux

export DEBIAN_DIR="/template/debian"
export PACKAGE="$INPUT_PACKAGE"
export MAINTAINER="$INPUT_MAINTAINER"
export VERSION="$INPUT_VERSION"
export ARCH="$INPUT_ARCH"
/replacetool

WORKDIR="/tmp/work"
PACKAGE_DIR="$WORKDIR/$PACKAGE"

mkdir -p /tmp/work/
cp -r /template "$PACKAGE_DIR"
cp -p "$PACKAGE" "$PACKAGE_DIR"
(
  cd "$PACKAGE_DIR"
  /git2chlog deb -o debian/changelog
  debian/rules build
  yes | debuild -us -uc
)
cp -p "$WORKDIR/"*.deb .

ls ./*.deb
