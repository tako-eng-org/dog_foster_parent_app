#!/bin/sh
echo "-----`basename $0`-----"

# エラー対策
# cgo: exec gcc: exec: "gcc": executable file not found in $PATH
export CGO_ENABLED=0

echo "-----dlv realize install start-----"
GO111MODULE=off go get -u github.com/derekparker/delve/cmd/dlv;
GO111MODULE=off go get -u github.com/oxequa/realize;
echo "-----dlv realize installed-----"

export GO111MODULE=on #go.modをgo build時に実行するための環境変数設定
export GOPATH=''

echo "-----remoteDebug or hotReload start-----"
#dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient
/go/bin/realize start;
echo "-----remoteDebug or hotReload end-----"
sleep 5000;
