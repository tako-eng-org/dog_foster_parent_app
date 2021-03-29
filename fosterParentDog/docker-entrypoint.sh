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
# golandでDockerコンテナをアタッチしてデバッグしたいとき
# port2345で待ち受けるのみ/サーバー(8090)を動かしたい場合は、golandの実行構成にあるのでデバッグボタンを押下すること。

# .realize.ymlのrunをfalseにする

# このスクリプトが実行されていることをログから確認したら、
# -----remoteDebug or hotReload start-----
# API server listening at: [::]:2345

# golandのデバッグボタン押下してサーバを起動する
# [GIN-debug] Listening and serving HTTP on :8090

# curlを叩いてブレークポイントを貼ったところで止まればOK
# curl -v localhost:8000/fosterparent/fetchAllRecords

# 1度実行すると不安定になるので、下記コマンドでコンテナ再起動した方が早いこともある(再現しない)
# docker-compose restart dog_app

#dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient
# ホットリロードでコンテナ起動したいとき
# realize.ymlにそってホットリロードを実行する
/go/bin/realize start;
echo "!!!!!!!!!exit!!!!!!!!!"
sleep 5000;
echo "-----remoteDebug or hotReload end-----"