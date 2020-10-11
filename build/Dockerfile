# syntax = docker/dockerfile:experimental
FROM golang:1.15.2 as build-env

RUN --mount=type=secret,id=credential,dst=/root/.gitconfig \
    go get github.com/narakosen-festival-info-2020/reversi-back/cmd/reversi_app && \ 
    CGO_ENABLED=0 go build -o /go/bin/reversi-back github.com/narakosen-festival-info-2020/reversi-back/cmd/reversi_app 

FROM alpine:3.12
COPY --from=build-env /go/bin/reversi-back .
ENTRYPOINT ["/reversi-back"]