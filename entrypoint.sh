#!/bin/sh

set -eux

INPUT_VERSION="$(echo "$INPUT_VERSION" | sed -E "s,^refs/tags/,,")"

if [ -z "$INPUT_INSTALLED_SIZE" ]; then
  PACKAGE_ROOT_SIZE_BYTES="$(du --bytes --summarize --exclude=DEBIAN "$INPUT_PACKAGE_ROOT"/ | awk '{print $1}')"
  INPUT_INSTALLED_SIZE="$(( (PACKAGE_ROOT_SIZE_BYTES + 1024 - 1) / 1024 ))"
fi

case "${INPUT_COMPRESS_TYPE}" in
  gzip | xz | zstd | none)
    # nothing to do
    ;;
  *)
    echo "[ERR] unsupported compress type. input is '${INPUT_COMPRESS_TYPE}'"
    exit 1
    ;;
esac

replacetool /replacetool_template/control /template/DEBIAN/control

cp -r /template/DEBIAN "$INPUT_PACKAGE_ROOT/"

# remove prefix 'v'
FIXED_VERSION="$(echo "$INPUT_VERSION" | sed -E 's/^v//')"
readonly FIXED_VERSION

OWNER="--root-owner-group"
if [ "${INPUT_KEEP_OWNERSHIP}" = "true" ]; then
  OWNER=""
fi

# create deb file
readonly DEB_FILE="${INPUT_PACKAGE}_${FIXED_VERSION}_${INPUT_ARCH}.deb"
dpkg-deb -Z"${INPUT_COMPRESS_TYPE}" ${OWNER:+"$OWNER"} --build "$INPUT_PACKAGE_ROOT" "$DEB_FILE"

ls ./*.deb
echo "file_name=$DEB_FILE" >> "${GITHUB_OUTPUT}"
