# syntax = docker/dockerfile:experimental
FROM golang:1.15.2 AS build-env

RUN --mount=type=secret,id=credential,dst=/root/.gitconfig \
    go get github.com/narakosen-festival-info-2020/reversi-back && \
    go build -o $GOPATH/bin/reversi-back github.com/narakosen-festival-info-2020/reversi-back/cmd/reversi_app

FROM busybox
COPY --from=build-env $GOPATH/bin/reversi-back /reversi-back
ENTRYPOINT [ "/reversi-back" ]