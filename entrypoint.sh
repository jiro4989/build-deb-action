#!/bin/sh

set -eux

INPUT_VERSION="$(echo "$INPUT_VERSION" | sed -E "s,^refs/tags/,,")"

if [ -z "$INPUT_INSTALLED_SIZE" ]; then
  PACKAGE_ROOT_SIZE_BYTES="$(du -bcs --exclude=DEBIAN "$INPUT_PACKAGE_ROOT"/ | awk '{print $1}' | head -1 | sed -e 's/^0\+//')"
  INPUT_INSTALLED_SIZE="$(( (PACKAGE_ROOT_SIZE_BYTES + 1024 - 1) / 1024 ))"
fi

case "${INPUT_COMPRESS_TYPE}" in
  gzip | xz | zstd)
    # nothing to do
    ;;
  *)
    echo "[ERR] unsupported compress type. input is '${INPUT_COMPRESS_TYPE}'"
    exit 1
    ;;
esac

/replacetool \
  --debian-dir:/template/DEBIAN \
  --package:"$INPUT_PACKAGE" \
  --version:"$INPUT_VERSION" \
  --installed-size:"$INPUT_INSTALLED_SIZE" \
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
dpkg-deb -Z"${INPUT_COMPRESS_TYPE}" -b "$INPUT_PACKAGE_ROOT" "$DEB_FILE"

ls ./*.deb
echo "file_name=$DEB_FILE" >> "${GITHUB_OUTPUT}"
