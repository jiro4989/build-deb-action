#!/bin/sh

set -eux

/replacetool \
  --debian-dir:/template/DEBIAN \
  --package:"$INPUT_PACKAGE" \
  --version:"$INPUT_VERSION" \
  --arch:"$INPUT_ARCH" \
  --maintainer:"$INPUT_MAINTAINER" \
  --description:"$INPUT_DESC"

cp -r /template/DEBIAN "$INPUT_PACKAGE_ROOT/"

# remove prefix 'v'
FIXED_VERSION="$(echo "$INPUT_VERSION" | sed -E 's/^v//')"
readonly FIXED_VERSION

# create deb file
readonly DEB_FILE="${INPUT_PACKAGE}_${FIXED_VERSION}_${INPUT_ARCH}.deb"
dpkg-deb -b "$INPUT_PACKAGE_ROOT" "$DEB_FILE"

ls ./*.deb
