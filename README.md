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
