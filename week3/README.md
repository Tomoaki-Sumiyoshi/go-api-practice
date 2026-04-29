# Week 3: net/httpでAPIサーバを作る

## 目的

Go標準ライブラリの `net/http` を使い、HTTP APIサーバを作る。  
JSONレスポンス、HTTPメソッド分岐、クエリパラメータ、パスパラメータ相当の処理、簡易CRUD APIの基礎を理解する。

## 使用教材

### 1. Go公式 Writing Web Applications

URL: https://go.dev/doc/articles/wiki/

学習範囲:

- Introduction
- Getting Started
- Data Structures
- Introducing the net/http package
- Using net/http to serve wiki pages
- Handling non-existent pages
- Error handling

### 2. Go公式 net/http package

URL: https://pkg.go.dev/net/http

学習範囲:

- Overview
- `HandleFunc`
- `ListenAndServe`
- `Request`
- `ResponseWriter`
- `ServeMux`

### 3. Go公式 Effective Go

URL: https://go.dev/doc/effective_go

学習範囲:

- Errors
- The blank identifier
- Embedding

## 理解する内容

- `http.HandleFunc`
- `http.ListenAndServe`
- `http.ResponseWriter`
- `*http.Request`
- JSONレスポンス
- HTTPステータスコード
- クエリパラメータ
- パスパラメータ相当の処理
- HTTPメソッドごとの分岐
- 簡易CRUD API

## 練習問題

### 01-hello-http

`GET /health` で `ok` を返すHTTPサーバを作る。

#### 仕様

- `localhost:8080` で起動する
- `GET /health` にアクセスすると `ok` を返す

#### 解けるようになる教材到達点

- Go公式 Writing Web Applications の Introducing the net/http package 完了後

#### 実行例

```bash
go run ./week3/01-hello-http
```

確認

```bash
curl http://localhost:8080/health
```

### 02-json-response

`GET /me` でJSONのUser情報を返す。

#### 仕様

レスポンス例:

```json
{
  "id": 1,
  "name": "Taro",
  "email": "taro@example.com"
}
```

#### 解けるようになる教材到達点

- Go公式 net/http package の `ResponseWriter`、`Request` 確認後

#### 実行例

```bash
go run ./week3/02-json-response
```

確認

```bash
curl http://localhost:8080/me
```

### 03-query-and-path

`GET /greet?name=Taro` でJSONレスポンスを返す。

#### 仕様

リクエスト

```bash
curl "http://localhost:8080/greet?name=Taro"
```

レスポンス

```json
{
  "message": "Hello, Taro"
}
```

`name` が空の場合:

```json
{
  "message": "Hello, Guest"
}
```

#### 解けるようになる教材到達点

- Go公式 net/http package の `Request` 確認後

#### 実行例

```bash
go run ./week3/03-query-and-path
```

### 04-method-routing

`/items` に対して `GET` と `POST` をメソッドで分岐する。

#### 仕様

- `GET /items`
  - item一覧を返す
- `POST /items`
  - itemを作成した想定のレスポンスを返す
- それ以外のメソッド
  - `405 Method Not Allowed` を返す

#### 解けるようになる教材到達点

- Go公式 net/http package の `HandleFunc`、`ServeMux` 確認後

#### 実行例

```bash
go run ./week3/04-method-routing
```

確認

```bash
curl http://localhost:8080/items
curl -X POST http://localhost:8080/items
curl -X DELETE http://localhost:8080/items
```

### 05-mini-crud-memory

インメモリでTodo APIを作る。

#### 仕様

作成するエンドポイント:

| Method | Path          | 内容         |
| ------ | ------------- | ------------ |
| GET    | `/todos`      | Todo一覧取得 |
| POST   | `/todos`      | Todo作成     |
| GET    | `/todos/{id}` | Todo詳細取得 |

Todoの構造:

```go
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```

#### 解けるようになる教材到達点

- Week3の全教材完了後

#### 実行例

```bash
go run ./week3/05-mini-crud-memory
```

確認

```bash
curl http://localhost:8080/todos
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go"}'
curl http://localhost:8080/todos/1
```

## フォルダ構成

```
week3/
├── 01-hello-http/
│   └── main.go
├── 02-json-response/
│   └── main.go
├── 03-query-and-path/
│   └── main.go
├── 04-method-routing/
│   └── main.go
├── 05-mini-crud-memory/
│   ├── main.go
│   ├── todo.go
│   └── store.go
└── README.md
```
