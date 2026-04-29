# Go API Practice

Next.jsエンジニアがGoを1か月学習し、標準ライブラリで簡易APIサーバを作るための練習リポジトリです。

## Goal

- Goの基本文法を理解する
- Go module / package / testを理解する
- net/httpでAPIサーバを作る
- 簡易CRUD APIを実装する

## Final API

| Method | Path        | Description  |
| ------ | ----------- | ------------ |
| GET    | /health     | health check |
| GET    | /todos      | list todos   |
| GET    | /todos/{id} | get todo     |
| POST   | /todos      | create todo  |
| PUT    | /todos/{id} | update todo  |
| DELETE | /todos/{id} | delete todo  |

## Learning Log

- Week 1: Go basics
- Week 2: modules, packages, tests, JSON
- Week 3: net/http basics
- Week 4: final API server
