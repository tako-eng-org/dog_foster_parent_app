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
|- db（DB操作にまつわる処理）
|   |- user_db.go
|   |- post_db.go
|   |- post_image_db.go
|   |_ 省略
|_ entity（テーブル定義/マイグレーション用）
    |- user.go
    |- post.go
    |- post_image.go
    |_ 省略
controllers
|- post_controller.go
|- user_controller.go
|_ 省略
serializers (エンドポイントごとに存在する。DB取得データをレスポンスに整形する)
|- index
|   |_ index_seliarizer.go
|- detail
|   |_ detail_serializer.go
|_ 省略
validators (エンドポイントごとに存在する。リクエストをDBデータ型に整形する)
|- post_edit
|   |_ post_edit_validator.go
|_ 省略
frontend（UIに関する資材）
|_ 省略
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
