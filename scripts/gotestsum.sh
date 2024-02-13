#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(dirname -- "${BASH_SOURCE[0]}")"
GOTESTSUM="$SCRIPT_DIR/../bin/gotestsum"


# We need to classify arguments into "go test" flags vs packages since gotestsum
# requires packages to be passed via a flag instead of positional args.
flag_args=()
packages=()

while [[ $# -gt 0 ]]; do
	if [[ $1 == -* ]]; then
		# If the flag starts with "-", treat it as a go test flag.
		flag_args+=("$1")
	else
		# Otherwise, it's a package to test.
		packages+=("$1")
	fi

	shift
done


set -x

# packages[*] is used to produce a single space-separated string for all packages
# flag_args[@] is used to produce separate quoted strings for each flag.
"$SCRIPT_DIR/../bin/gotestsum" --format=standard-json --rerun-fails --packages "${packages[*]:-}" -- "${flag_args[@]:-}"
