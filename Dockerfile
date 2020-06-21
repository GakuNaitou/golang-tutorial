FROM golang:1.11.2-alpine3.8 AS build

WORKDIR /go
ADD . /go

# TODO: packageのインストールができていないので修正する
RUN apk update && \
    apk add --no-cache git && \
    go get github.com/go-sql-driver/mysql && \  
    go get github.com/labstack/echo/middleware && \
    go get github.com/jinzhu/gorm

# 上記のpackageのインストールが上手くできていないため、main.goを走らせてもエラーになってしまうためコメントアウト
# CMD ["go", "run", "main.go"]