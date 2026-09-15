# Go API

Goであそぶ

## API Design

### API Category

このAPIはPublic APIです。

`/api/v1/*` へのアクセスにはAPIキーが必要です。APIキーは `X-API-Key` ヘッダーに指定します。

`/health` は認証なしでアクセスできます。

```json
{
  "meta": {
    "api": {
      "name": "go-api",
      "language": "Go",
      "category": "public"
    }
  }
}
```

### Environment Variables

```text
API_KEY=your-api-key
PORT=3000
```

`PORT` が指定されていない場合は `3000` を使用します。

### Endpoints

#### `GET /health`

APIの稼働状態を確認します。

#### `GET /api/v1/posts`

投稿一覧を取得します。

`username` クエリパラメータによる絞り込みにも対応しています。

```text
GET /api/v1/posts?username=hamru
```

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

### Response

```json
{
  "meta": {
    "api": {
      "name": "go-api",
      "language": "Go",
      "category": "public"
    },
    "count": 1
  },
  "data": {
    "posts": [
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
    ]
  }
}
```

### Error Response

```json
{
  "error": {
    "code": "NOT_FOUND",
    "message": "Post not found"
  }
}
```

主なHTTPステータス:

- `200 OK` - リクエスト成功
- `400 Bad Request` - リクエストが不正
- `401 Unauthorized` - APIキーが無効
- `404 Not Found` - リソースが存在しない

## Data Source

JSONファイルを使用します。

## Setup

```bash
API_KEY=test-api-key go run ./cmd/api
```

## Docker

```bash
docker build -t go-api .
```

```bash
docker run --rm \
  -p 3000:3000 \
  -e API_KEY=test-api-key \
  go-api
```

## License

このリポジトリは学習・技術検証目的で公開しています。

著作権は作者に帰属します。
無断転載・再配布・商用利用はご遠慮ください。

This repository is published for learning and technical verification purposes.

All rights to the content belong to the author.

Please do not reproduce, redistribute, or use any part of this project for commercial purposes without permission.

## Author

- h-waji (hamltail)
