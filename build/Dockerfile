# syntax = docker/dockerfile:experimental
FROM golang:1.15.2 as dev-builder

WORKDIR /go/src/github.com/narakosen-festival-info-2020/reversi-back/
COPY ./* ./
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=0 go build -o /go/bin/reversi-back ./cmd/reversi_app/


FROM golang:1.15.2 as main-builder

RUN --mount=type=secret,id=credential,dst=/root/.gitconfig \
    go get github.com/narakosen-festival-info-2020/reversi-back/cmd/reversi_app && \ 
    CGO_ENABLED=0 go build -o /go/bin/reversi-back github.com/narakosen-festival-info-2020/reversi-back/cmd/reversi_app 


FROM alpine:3.12 as develop
COPY --from=dev-builder /go/bin/reversi-back .
ENTRYPOINT [ "./reversi-back" ]


FROM alpine:3.12 as main
COPY --from=main-builder /go/bin/reversi-back .
ENTRYPOINT [ "./reversi-back" ]
