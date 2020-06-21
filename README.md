# ページ

### 登録周り
/login
/signup

### 投稿周り
/posts
/posts/:id
/posts/:id/edit

### ユーザー情報
/user

# 機能
### 認証周り
- 名前とメールアドレスの組み合わせによる認証
- ログイン時にcookie(userId)を付与
- ログアウト時と退会時にcookieを削除

### 投稿一覧
- 投稿一覧画面は誰でも見れる
- 新しい順に上から投稿(投稿者と内容)が表示される
- スレッドをクリックするとその投稿に対するスレッドを見れる
- 投稿を行える
- 自分が投稿したものは編集、削除ができる
- ログアウトもこのページから行える

### 投稿詳細(スレッド)
- ログインしていないと見れない
- 親となる投稿が一番上に表示され、その下にスレッドの投稿が新しい順で上から表示される
- スレッドの詳細からそのスレッドに対するスレッドを見れる
- スレッドの投稿を行える
- 自分が投稿したものは編集、削除ができる

# DB
- main.goの起動時にマイグレーションが行われ、指定したmodelの定義されている部分を元にDBの更新が行われる
- データの削除はdeleted_atを使ってソフトデリートしている

### ユーザー
- 名前とメールアドレスの情報を変更できる
- アカウントの削除(退会)もこのページから行える

# 環境構築
### localにプロジェクトをcloneする
```
$ git clone git@github.com:GakuNaitou/golang-tutorial.git
```
-> カレントディレクトリにプロジェクトが追加される

### プロジェクトディレクトリに移動する
```
$ cd golang-tutorial
```
-> プロジェクトディレクトリに移動

### Dockerコンテナをビルドする
```
$ docker-compose up --build
```
-> 特にエラーが出なければOK

### 別タブでアプリケーションサーバーにアクセスする
```
$ docker-compose exec app ash
```
-> アプリケーションサーバーに入る

### 必要なパッケージをインストールする(初回のみ)
```
[app-server]$ go get github.com/go-sql-driver/mysql
[app-server]$ go get github.com/labstack/echo/middleware
[app-server]$ go get github.com/jinzhu/gorm
```

### main.goを呼び出してアプリケーションを起動する
```
[app-server]$ go run main.go
```
-> 下記のような感じで表示されればOK
```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.1.16
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

### localhost:1323/loginにアクセスしてログイン画面が表示されれば完了


# DBサーバー,DBに入る方法
### DBサーバーに入る
```
$ docker exec -it go_db /bin/bash
```
-> DBサーバーに入る

### DB(MySql)に入る
```
[db-server]$ mysql -u db_user -h localhost -p
```
-> パスワードを求められるので下記を入力する
```
db_user_pwd
```
-> DB(MySql)に入る

### DBの確認
```
[mysql]$ show databases;
```
-> tutorial_db があればOK
