FROM golang:latest AS builder
ADD . /go/api
WORKDIR /go/api

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN rm -rf deploy
RUN mkdir deploy
RUN swag init -d adapter/http --parseDependency --parseInternal --parseDepth 3 -o adapter/http/rest/docs
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o goapp adapter/http/main.go
RUN mv goapp ./deploy/goapp
RUN mv adapter/http/rest/docs ./deploy/docs
RUN mv config.json ./deploy/config.json
RUN mv database ./deploy/database


FROM alpine:3.7 AS production
COPY --from=builder /go/api/deploy /api/

WORKDIR /api
ENTRYPOINT  ./goapp