FROM golang:1.22.1-alpine3.18

WORKDIR /app/

COPY ./back .

RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev && \
    go get -u github.com/gin-gonic/gin && \
    go get -u gorm.io/gorm

RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest

CMD ["air","-c",".air.toml"]

