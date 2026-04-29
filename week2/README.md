# Week 2: モジュール・パッケージ分割・テスト・JSON

## 目的

Go module、package分割、テスト、JSON処理を理解する。  
APIサーバの実装で使うドメイン処理を、HTTP層から独立した形で作れるようにする。

## 使用教材

### 1. Go公式 Tutorial: Create a Go module

URL: https://go.dev/doc/tutorial/create-module

学習範囲:

- ページ全体
- モジュール作成
- パッケージ作成
- 別モジュールからの利用

### 2. Go公式 Tutorial: Add a test

URL: https://go.dev/doc/tutorial/add-a-test

学習範囲:

- ページ全体
- Goの組み込みテスト機能

### 3. Go公式 testing package

URL: https://pkg.go.dev/testing

学習範囲:

- Overview
- Subtests and Sub-benchmarks

### 4. Go公式 Wiki: TableDrivenTests

URL: https://go.dev/wiki/TableDrivenTests

学習範囲:

- ページ全体

### 5. Go公式 How to Write Go Code

URL: https://go.dev/doc/code

学習範囲:

- Code organization

### 6. Go公式 Organizing a Go module

URL: https://go.dev/doc/modules/layout

学習範囲:

- ページ全体

## 理解する内容

- `go mod init`
- `go test`
- package分割
- public/private識別子
- table-driven test
- `encoding/json`
- ファイル入出力
- APIサーバで使うドメインモデル設計

## 練習問題

### 01-package-split

`calculator` packageを作り、`main.go` から呼び出す。

#### 仕様

- `calculator.Add(a, b int) int`
- `calculator.Sub(a, b int) int`
- `calculator.Mul(a, b int) int`
- `calculator.Div(a, b int) (int, error)`

#### 解けるようになる教材到達点

- Go公式 Tutorial: Create a Go module 完了後

#### 実行例

```bash
go run ./week2/01-package-split
```

### 02-table-test

Week1の以下の関数にtable-driven testを書く。

- Add
- IsEven
- ValidateAge

#### 解けるようになる教材到達点

- Go公式 Tutorial: Add a test 完了後
- Go公式 Wiki: TableDrivenTests 完了後

#### 実行例

```bash
go test ./week2/02-table-test
```

### 03-json-encode-decode

`User` structをJSON文字列に変換し、JSON文字列からstructへ戻す。

#### 仕様

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

以下を実装する。

- structからJSON文字列への変換
- JSON文字列からstructへの変換

#### 解けるようになる教材到達点

- Go公式 How to Write Go Code の Code organization 確認後
- encoding/json の基本確認後

#### 実行例

```bash
go run ./week2/03-json-encode-decode
```

### 04-file-io

JSONファイルに `[]User` を保存し、読み込む。

#### 仕様

- `users.json` に `[]User` を保存する
- `users.json` から `[]User` を読み込む
- 保存と読み込みの処理を関数に分ける

#### 解けるようになる教材到達点

- `encoding/json` とpackage分割練習後

#### 実行例

```bash
go run ./week2/04-file-io
```

### 05-todo-domain

HTTP APIで使う前段階として、Todoのドメイン処理を作る。

#### 仕様

`Todo` structを作る。

```go
type Todo struct {
    ID        int
    Title     string
    Completed bool
}
```

`TodoStore` を作り、以下を実装する。

- `Create(title string) Todo`
- `FindAll() []Todo`
- `FindByID(id int) (Todo, bool)`
- `Update(id int, title string, completed bool) (Todo, bool)`
- `Delete(id int) bool`

#### 解けるようになる教材到達点

- Week2の全教材完了後

#### 実行例

```bash
go test ./week2/05-todo-domain
```

## フォルダ構成

```
week2/
├── 01-package-split/
│   ├── main.go
│   └── calculator/
│       └── calculator.go
├── 02-table-test/
│   ├── calc.go
│   └── calc_test.go
├── 03-json-encode-decode/
│   └── main.go
├── 04-file-io/
│   ├── main.go
│   └── users.json
├── 05-todo-domain/
│   ├── todo.go
│   ├── store.go
│   └── store_test.go
└── README.md
```
