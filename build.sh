#!/bin/bash
set -euf -o pipefail

GOVVVBIN=$(which govvv) || GOVVVBIN=${GOVVVBIN:=$GOPATH/bin/govvv}
GOBIN=$(which go) || GOBIN=${GOBIN:=$GOROOT/bin/go}
ECHOBIN=$(which echo)
SEDBIN=$(which sed)

GOPKGPATH=$($GOBIN list ./cmd)
GOLDFLAGS=$($GOVVVBIN -flags)
GOLDFLAGS=$($ECHOBIN $GOLDFLAGS | $SEDBIN -e 's|main|'$GOPKGPATH'|g')

$ECHOBIN "Building..."
$GOBIN build -ldflags="$GOLDFLAGS" -o bin/paperless-cli
$ECHOBIN "Build complete."
