#!/bin/bash

# This script is used to automate incrementing Git tags. If the last tag is "v0.n.m" this script
# will create a new tag "v0.n.(m+1)" pointing to the latest commit. This script is derived, with
# minor modifications, from https://stackoverflow.com/questions/3760086/automatic-tagging-of-releases.

set -e

dry_run=false
for arg in "$@"; do
  case $arg in
  --dry-run | -d)
    dry_run=true
    shift # Remove --dry-run or -d from the arguments
    ;;
  esac
done

# As a precaution, we only allow tags to be created on clean branches to ensure that any changes
# the user wants to be included in the tag have been committed.
if [ -n "$(git status --porcelain)" ]; then
  echo "Can only add a tag from a clean branch, check git status and try again."
  exit 1
fi

NEXT_VERSION_SCRIPT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/next_version.sh"
NEW_VERSION=$(bash $NEXT_VERSION_SCRIPT --patch)

# Get the current hash and see if it already has a tag. We will only create the new tag if the
# current hash does not already have a tag.
GIT_COMMIT=$(git rev-parse HEAD)
NEEDS_TAG=$(git describe --contains "$GIT_COMMIT" 2>/dev/null || echo "")

if [ -n "$NEEDS_TAG" ]; then
  echo "Can only add a new tag to a commit which does not have one, but the latest commit already does."
  exit 1
fi

if [ "$dry_run" = true ]; then
  echo "Dry run mode is enabled. Skipping tag creation and push."
  echo "Would have updated $VERSION to $NEW_VERSION" by running:
  echo "     git tag $NEW_VERSION"
  echo "     git push origin $NEW_VERSION"
  exit 0
fi

echo "Updating $VERSION to $NEW_VERSION"
git tag $NEW_VERSION
echo "Pushing new tag $NEW_VERSION"
git push origin $NEW_VERSION
