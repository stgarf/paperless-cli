#!/bin/bash
set -euf -o pipefail

GOBIN=$(command -v go) || GOBIN=${GOBIN:=$GOROOT/bin/go}
ECHOBIN=$(command -v echo)
command -v govvv >/dev/null 2>&1 || {
    $ECHOBIN -en 'Need govvv, building and installing..';
    $GOBIN install github.com/ahmetb/govvv;
    $ECHOBIN '. Done.';
} 
GOVVVBIN=$(command -v govvv) || GOVVVBIN=${GOVVVBIN:=$GOPATH/bin/govvv}
SEDBIN=$(command -v sed)

GOPKGPATH=$($GOBIN list ./cmd)
GOLDFLAGS=$($GOVVVBIN -flags)
GOLDFLAGS=$($ECHOBIN $GOLDFLAGS | $SEDBIN -e 's|main|'$GOPKGPATH'|g')

$ECHOBIN -ne "Building.."
$GOBIN build -ldflags="$GOLDFLAGS" -o bin/paperless-cli
$ECHOBIN ". Build complete."
