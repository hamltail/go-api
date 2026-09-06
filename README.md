# Go API

GoであそぶAPI

## API Design

### Endpoints

#### `GET /health`

APIの稼働状態を確認します。

#### `GET /api/v1/posts`

投稿一覧を取得します。

#### `GET /api/v1/posts/:id`

指定したIDの投稿を取得します。

### Post

投稿データはPawthのデータ構造を参考にします。

```json
{
  "id": 1,
  "user": {
    "username": "hamru",
    "displayName": "はむる"
  },
  "content": "GoでAPIを作っています。",
  "postedOn": "2026-09-06",
  "createdAt": "2026-09-06T10:00:00+09:00"
}
```

## Data Source

初期実装ではJSONファイルを使用します。

## License

このリポジトリは学習・技術検証目的で公開しています。

著作権は作者に帰属します。
無断転載・再配布・商用利用はご遠慮ください。

This repository is published for learning and technical verification purposes.

All rights to the content belong to the author.

Please do not reproduce, redistribute, or use any part of this project for commercial purposes without permission.

## Author

- h-waji (hamltail)
