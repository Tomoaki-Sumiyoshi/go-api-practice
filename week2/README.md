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

```

```
