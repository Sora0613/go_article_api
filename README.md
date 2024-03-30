# Article Manager API

## Description
This API is used to manage articles. 

It allows to create and read articles.

## Source Codes
The source codes are available at ./back

## TODO
- 選考プロセスのstepsの上手な追加、管理方法
- 選考プロセスの順番の管理
- POST Articleの実装（あったら楽かも）

## The Content of tasks
テーマ: 選考体験記の投稿・取得ができるAPIを実装する

- 必須技術:
  - [x] Go言語 (Gin + Gorm), MySQL
  - [x] Docker
  - [x] GitHub

- 必須課題
  - [x] GitHubでコードを管理している
  - [x] Dockerを使用してローカルの環境構築ができる
  - [x] データがDBに保存され、取得されるようになっている

- 発展課題 
  - [ ] OpenAPI(Swagger)を使ったAPIドキュメントの作成
  - [ ] データ量が大量になっても素早く動くようにパフォーマンスが考慮されている 
  - [ ] コードの簡潔さや再利用可能性を考慮したGoの実装 
  - UXを考えた設計
    - [x] 任意の面接数に対応できている 
    - [ ] 任意の質問項目に対応できている 
    - [ ] 質問内容を後からDB操作で変えることで対応できる