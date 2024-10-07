#!/bin/bash

set -e

# This script is used to generate a new release tag that is 1 minor or patch version greater than the most
# recent release. Adapted from https://stackoverflow.com/questions/3760086/automatic-tagging-of-releases.
patch_version=false
for arg in "$@"; do
    case $arg in
    --patch | -p)
        patch_version=true
        shift # Remove --patch or -p from the arguments
        ;;
    esac
done

VERSION=$(git describe --abbrev=0 --tags)

# Remove leading "v" from version tag (for details, see the "Substring Extraction" section of
# https://www.tldp.org/LDP/abs/html/string-manipulation.html).
VERSION="${VERSION:1}"

# Replace "." with a space so we can split the version into an array.
read -r -a VERSION_BITS <<<"${VERSION//./ }"

MAJOR_VERSION=${VERSION_BITS[0]}
MINOR_VERSION=${VERSION_BITS[1]}
PATCH_VERSION=${VERSION_BITS[2]:0}

# Increment the minor version if --patch is not set
if [ "$patch_version" = false ]; then
    MINOR_VERSION=$((MINOR_VERSION + 1))
else
    PATCH_VERSION=$((PATCH_VERSION + 1))
fi

NEW_VERSION="v${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}"

echo ${NEW_VERSION}
