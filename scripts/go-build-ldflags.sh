#!/bin/sh

snapshot=false
for arg in "$@"; do
    case $arg in
    --snapshot | -p)
        snapshot=true
        shift # Remove --patch or -p from the arguments
        ;;
    esac
done

if [ "$snapshot" = false ]; then
    export GIT_VERSION=$(./scripts/next_version.sh || echo unknown)
else
    export GIT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo unknown)
fi

export GIT_REVISION=$(git rev-parse --short HEAD)
export BUILD_DATE=$(date '+%F-%T') # outputs something in this format 2017-08-21-18:58:45
export BASE_PACKAGE=github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/buildinfo

LD_FLAGS="-X ${BASE_PACKAGE}.SHA=${GIT_REVISION} \
-X ${BASE_PACKAGE}.Version=${GIT_VERSION} \
-X ${BASE_PACKAGE}.Date=${BUILD_DATE}"

echo $LD_FLAGS
