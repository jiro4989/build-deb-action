#!/bin/sh

set -eux

/replacetool \
  --debian-dir:/template/debian \
  --package:"$INPUT_PACKAGE" \
  --maintainer:"$INPUT_MAINTAINER" \
  --version:"$INPUT_VERSION" \
  --arch:"$INPUT_ARCH"

WORKDIR="/tmp/work"
PACKAGE_DIR="$WORKDIR/$INPUT_PACKAGE"

mkdir -p /tmp/work/
cp -r /template "$PACKAGE_DIR"
cp -p "$INPUT_PACKAGE" "$PACKAGE_DIR"
(
  cd "$PACKAGE_DIR"
  /git2chlog deb -o debian/changelog
  debian/rules build
  yes | debuild -us -uc
)
cp -p "$WORKDIR/"*.deb .

ls ./*.deb
