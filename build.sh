#!/bin/bash
set -euf -o pipefail

GOBIN=$(command -v go) || GOBIN=${GOBIN:=$GOROOT/bin/go}
ECHOBIN=$(command -v echo)
command -v govvv >/dev/null 2>&1 || {
    $ECHOBIN -en 'Need govvv, building and installing..';
    $GOBIN install github.com/ahmetb/govvv;
    $ECHOBIN '. Done.';
} 
command -v gox >/dev/null 2>&1 || {
    $ECHOBIN -en 'Need gox, building and installing..';
    $GOBIN install github.com/mitchellh/gox;
    $ECHOBIN '. Done.';
} 
GOVVVBIN=$(command -v govvv) || GOVVVBIN=${GOVVVBIN:=$GOPATH/bin/govvv}
GOXBIN=$(command -v gox) || GOXBIN=${GOXBIN:=$GOPATH/bin/gox}

$ECHOBIN -ne "Building.."
$GOXBIN -os="linux darwin windows" -arch="amd64" -output="./bin/paperless-cli.{{.OS}}.{{.Arch}}" -ldflags "$($GOVVVBIN -flags -pkg $($GOVVVBIN list ./cmd))" -verbose ./...
$ECHOBIN ". Build complete."
