# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# boiler-plateディレクトリの作成
# RUN mkdir /go/src/github.com/boiler-plate
# ワーキングディレクトリの設定
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY . .
# パッケージのインポート
# RUN go get -u golang.org/x/tools/cmd/goimports　だとできない（gogetが使えなくなった？）
RUN go install golang.org/x/tools/cmd/goimports@latest
# アップデートとgitのインストール
RUN apk update && apk add git
# ポート開放
EXPOSE 8080
CMD ["go", "run", "main.go"]