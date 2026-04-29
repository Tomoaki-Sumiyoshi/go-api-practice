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
