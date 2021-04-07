## 犬の里親募集アプリ

### start
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

### ディレクトリ構成(編集中)
参考元(gin公式サンプル)：
[https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/serializers.go#L14-L34](https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/serializers.go#L14-L34)

```
models
|- db.go
|   
|_ entities
    |- user.go
    |- post.go
    |_ post_image.go
    |_省略
controllers
|- post_controller.go
|- user_controller.go
|_省略
serializers (controllerと対)
|- post_serializer.go
|- user_serializer.go
|_省略
```

### GoLand(IDE)でDockerコンテナへリモートデバッグする方法
```docker-compose logs -f dog_app```
でログを確認する。<br>
下記ログ出力していればdlvは起動している。
```
-----remoteDebug or hotReload start-----
API server listening at: [::]:2345
```
golandのデバッグ実行ボタン押下し、アプリケーションサーバを起動する。
下記ログ出力していればアプリケーションサーバは起動している。

```
[GIN-debug] Listening and serving HTTP on :8090
```
curlなどでAPIを叩いてみて、ブレークポイントを貼ったところで止まればリモートデバッグ機能は期待動作している。<br>
`curl -v localhost:8000/fosterparent/fetchAllRecords`等<br>
<br>
注意点
1度実行すると不安定になるので、下記コマンドでコンテナ再起動した方が早いこともある(再現しない)<br>
`docker-compose restart dog_app`<br>

# ホットリロードを使用したい場合
`fosterParentDog/.realize.yaml`にて設定すること。<br>
