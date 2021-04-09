#!/bin/sh
echo "-----`basename $0`-----"

export GO111MODULE=on #go.modをgo build時に実行するための環境変数設定
export GOPATH=''
export CGO_ENABLED=0

echo "-----remoteDebug or hotReload start-----"
#dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient
/go/bin/realize start;
echo "-----remoteDebug or hotReload end-----"
sleep 5000;
exit 1