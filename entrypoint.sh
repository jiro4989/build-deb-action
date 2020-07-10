#!/bin/sh

set -eux

/replacetool \
  --debian-dir:/template/DEBIAN \
  --package:"$INPUT_PACKAGE" \
  --version:"$INPUT_VERSION" \
  --arch:"$INPUT_ARCH" \
  --maintainer:"$INPUT_MAINTAINER" \
  --description:"$INPUT_DESC"

readonly WORKDIR="/tmp/work"
readonly PACKAGE_DIR="$WORKDIR/$INPUT_PACKAGE"

mkdir -p /tmp/work/
cp -r /template "$PACKAGE_DIR"

readonly INSTALL_DIR="$PACKAGE_DIR/usr/bin/"
mkdir -p "$INSTALL_DIR"
install -m 0755 "$INPUT_PACKAGE" "$INSTALL_DIR"

# remove prefix 'v'
FIXED_VERSION="$(echo "$INPUT_VERSION" | sed -E 's/^/v/')"
readonly FIXED_VERSION

# create deb file
readonly DEB_FILE="${INPUT_PACKAGE}_${FIXED_VERSION}_${INPUT_ARCH}.deb"
dpkg-deb -b "$PACKAGE_DIR" "$DEB_FILE"

cp -p "$WORKDIR/"*.deb .

ls ./*.deb
