# dog_foster_parent_app
犬の里親募集アプリ

```sh
# ローカル開発する場合に必要なモジュール
go get realize

# ビルド
docker-compose build --no-cache

# アプリ起動(nginxを経由するアプリは全て起動します)
docker-compose up -d nginx
# access to http://localhost:8000
# 各コンテナへは、nginx経由でのみアクセス可能です

# pgAdmin
docker-compose up -d pgadmin
# access to http://localhost:8001
```


編集中
```
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

```
