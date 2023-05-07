# reportify backend

![Deploy](https://github.com/fy23-gw-gackathon/reportify-backend/workflows/Deploy/badge.svg)
![version](https://img.shields.io/badge/version-1.0--SNAPSHOT-blue.svg)

## Description
#### ツールインストール
go install github.com/cosmtrek/air@v1.43.0 \
&& go install github.com/swaggo/swag/cmd/swag@latest

#### ローカル開発の場合はディレクトリ内に.envを作成
```shell
APP_DEBUG=true
```

#### コード生成時は以下のコマンドを実行
```shell
make gen
```
