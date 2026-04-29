# Week 3: net/httpでAPIサーバを作る

## 目的

Go標準ライブラリの `net/http` を使い、HTTP APIサーバを作る。  
JSONレスポンス、HTTPメソッド分岐、クエリパラメータ、簡易CRUD APIの基礎を理解する。

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
