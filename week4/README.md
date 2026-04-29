# Week 4: 最終APIサーバとしてまとめる

## 目的

Week1からWeek3で学習した内容を使い、GitHub公開できる簡易CRUD APIサーバを作る。  
handler、model、store、responseに分け、テストを含む構成にする。

## 使用教材

### 1. Go公式 Organizing a Go module

URL: https://go.dev/doc/modules/layout

学習範囲:

- Basic package
- Command package
- Package or command with supporting packages

### 2. Go公式 How to Write Go Code

URL: https://go.dev/doc/code

学習範囲:

- Code organization

### 3. Go公式 testing package

URL: https://pkg.go.dev/testing

学習範囲:

- Examples
- Subtests and Sub-benchmarks

### 4. Go公式 net/http package

URL: https://pkg.go.dev/net/http

学習範囲:

- `Handler`
- `HandlerFunc`
- `ServeMux`
- `ResponseWriter`
- `Request`

## 理解する内容

- APIサーバのフォルダ分割
- handler / store / model の分離
- バリデーション
- エラーレスポンス
- handlerのテスト
- README整備
- GitHub公開用の最終整理

## 練習問題

### 01-layered-api

`handler`、`model`、`store` に分けて `GET /todos` を作る。

#### 仕様

- `model.Todo` を定義する
- `store.MemoryStore` を定義する
- `handler.TodoHandler` を定義する
- `GET /todos` でTodo一覧を返す

#### 解けるようになる教材到達点

- Go公式 Organizing a Go module 完了後

#### 実行例

```bash
go run ./week4/01-layered-api
```

### 02-validation-errors

`POST /todos` で空titleなら `400 Bad Request` を返す。

#### 仕様

正常リクエスト:

```json
{
  "title": "Learn Go"
}
```

異常リクエスト:

```json
{
  "title": ""
}
```

異常時レスポンス:

```json
{
  "error": "title is required"
}
```

#### 解けるようになる教材到達点

- Go公式 net/http package の `ResponseWriter`、ステータスコード確認後

#### 実行例

```bash
go run ./week4/02-validation-errors
```

### 03-handler-tests

`httptest` を使ってhandlerのテストを書く。

#### 仕様

以下のテストを書く。

- `GET /health` が `200 OK` を返す
- `POST /todos` が正常リクエストで `201 Created` を返す
- `POST /todos` が空titleで `400 Bad Request` を返す

#### 解けるようになる教材到達点

- Go公式 testing package 確認後
- Go公式 net/http package 確認後

#### 実行例

```bash
go test ./week4/03-handler-tests/...
```

### 04-final-api-server

簡易CRUD APIサーバを完成させる。

#### 仕様

作成するエンドポイント:
| Method | Path | 内容 |
| ------ | ------------- | -------- |
| GET | `/health` | ヘルスチェック |
| GET | `/todos` | Todo一覧取得 |
| GET | `/todos/{id}` | Todo詳細取得 |
| POST | `/todos` | Todo作成 |
| PUT | `/todos/{id}` | Todo更新 |
| DELETE | `/todos/{id}` | Todo削除 |

Todoの構造:

```go
type Todo struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}
```

エラーレスポンス:

```json
{
  "error": "message"
}
```

#### 解けるようになる教材到達点

- Week1からWeek4の全教材完了後

#### 実行例

```bash
go run ./week4/04-final-api-server/cmd/api
```

確認

```bash
curl http://localhost:8080/health

curl http://localhost:8080/todos

curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go"}'

curl http://localhost:8080/todos/1

curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go deeply","completed":true}'

curl -X DELETE http://localhost:8080/todos/1
```

#### 最終APIサーバ構成

```
week4/04-final-api-server/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── todo_handler.go
│   │   └── todo_handler_test.go
│   ├── model/
│   │   └── todo.go
│   ├── store/
│   │   ├── memory_store.go
│   │   └── memory_store_test.go
│   └── response/
│       └── json.go
├── README.md
└── go.mod
```

## フォルダ構成

```
week4/
├── 01-layered-api/
├── 02-validation-errors/
├── 03-handler-tests/
├── 04-final-api-server/
└── README.md
```
