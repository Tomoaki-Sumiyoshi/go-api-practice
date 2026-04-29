# Week 1: Goの文法・型・関数・エラー処理に慣れる

## 目的

Goの基本文法に慣れ、CLIプログラムを書けるようにする。  
APIサーバ作成の前提になる、型、関数、slice、map、struct、method、error処理を理解する。

## 使用教材

### 1. Go公式 Tutorial: Get started with Go

URL: https://go.dev/doc/tutorial/getting-started

学習範囲:

- ページ全体
- Hello, World
- `go run`
- 外部モジュール呼び出し

### 2. Go公式 A Tour of Go

URL: https://go.dev/tour/

学習範囲:

- Basics
- Flow control statements
- More types

### 3. Go公式 Effective Go

URL: https://go.dev/doc/effective_go

学習範囲:

- Formatting
- Names
- Control structures
- Functions
- Data
- Methods
- Errors

## 理解する内容

- `go run`
- `package main`
- `func main`
- 変数
- 定数
- 型
- `if`
- `for`
- `switch`
- 配列
- slice
- map
- struct
- method
- pointerの基本
- `error` の扱い
- `fmt` パッケージの基本

## 練習問題

### 01-hello-cli

コマンドラインに `"Hello, Go"` を出力する。

#### 解けるようになる教材到達点

- Go公式 Tutorial: Get started with Go 完了後

#### 実行例

```bash
go run ./week1/01-hello-cli
```

### 02-types-and-functions

以下の関数を実装する。

- Add(a, b int) int
- IsEven(n int) bool
- Max(a, b int) int

#### 解けるようになる教材到達点

- A Tour of Go の Basics 完了後

#### 実行例

```bash
go run ./week1/02-types-and-functions
```

### 03-slices-and-maps

文字列sliceを受け取り、出現回数を map[string]int で返す。

例

```go
[]string{"go", "js", "go"}
```

期待する結果

```go
map[string]int{
    "go": 2,
    "js": 1,
}
```

#### 解けるようになる教材到達点

- A Tour of Go の More types 完了後

#### 実行例

```bash
go run ./week1/03-slices-and-maps
```

### 04-struct-method

User structを作り、FullName() メソッドを実装する。

例

```go
type User struct {
    FirstName string
    LastName  string
}
```

期待する結果

```go
user := User{FirstName: "Taro", LastName: "Yamada"}
fmt.Println(user.FullName())
// Taro Yamada
```

#### 解けるようになる教材到達点

- Effective Go の Data、Methods 完了後

#### 実行例

```bash
go run ./week1/04-struct-method
```

### 05-error-handling

年齢を受け取り、0未満ならerrorを返す ValidateAge を作る。

#### 仕様

- age >= 0 の場合は nil を返す
- age < 0 の場合は error を返す

#### 解けるようになる教材到達点

- Effective Go の Errors 完了後

#### 実行例

```bash
go run ./week1/05-error-handling
```

## フォルダ構成

```
week1/
├── 01-hello-cli/
│   └── main.go
├── 02-types-and-functions/
│   └── main.go
├── 03-slices-and-maps/
│   └── main.go
├── 04-struct-method/
│   └── main.go
├── 05-error-handling/
│   └── main.go
└── README.md
```
