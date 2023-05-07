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
OPENAI_KEY=<slackで確認 chatGPTのトークン>
```

#### コード生成時は以下のコマンドを実行
```shell
make gen
```

#### ローカルで日報作成時にはローカル開発用の.envを作成して、`/organizations/{organizationCode}/reports`に以下のクエリを投げてみる
組織コード:NewGraduateTraining2
日報作成リクエスト:
```
{
  "body": "# 日報 2023-05-02\n\n## 今日やったこと\n- インバスケット研修\n- Go研修\n\n## 学んだこと、感じたこと\n- インバスケット研修でリアルの課題の解決に挑戦しましたが、限られた時間で問題を正確に解ききることはできず、自分自身はまだ実際のトラベルへの対応力は圧倒的に足りないことに気づきました\n- 元々Goの経験はあったのですが、チームワークではgoroutineに詰まってしまったので、本質的な部分の理解はまだ足りないことに気づきました\n\n## 明日やること\n- MySQL研修\n- CS研修\n- 自己理解研修\n",
  "tasks": []
}
```