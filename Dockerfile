FROM golang:latest AS builder
ADD . /go/api
WORKDIR /go/api

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN rm -rf deploy
RUN mkdir deploy
RUN swag init -d adapter/http --parseDependency --parseInternal --parseDepth 2 -o adapter/http/docs
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o goapp adapter/http/main.go
RUN mv goapp ./deploy/goapp
RUN mv adapter/http/docs ./deploy/docs
RUN mv config.json ./deploy/config.json


FROM alpine:3.7 AS production
COPY --from=builder /go/api/deploy /api/

WORKDIR /api
ENTRYPOINT  ./goapp