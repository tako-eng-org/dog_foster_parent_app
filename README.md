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
