#!/bin/sh

set -eux

INPUT_VERSION="$(echo "$INPUT_VERSION" | sed -E "s,^refs/tags/,,")"

if [ -z "$INPUT_SIZE" ]; then
  PACKAGE_ROOT_SIZE_BYTES="$(du -bcs --exclude=DEBIAN "$INPUT_PACKAGE_ROOT"/ | awk '{print $1}' | head -1 | sed -e 's/^0\+//')"
  INPUT_SIZE="$(awk "BEGIN {print ($PACKAGE_ROOT_SIZE_BYTES/1024)+1}" | awk '{print int($0)}')"
fi

/replacetool \
  --debian-dir:/template/DEBIAN \
  --package:"$INPUT_PACKAGE" \
  --version:"$INPUT_VERSION" \
  --size:"$INPUT_SIZE" \
  --depends:"$INPUT_DEPENDS" \
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
