FROM golang:1.21.5-alpine3.19

WORKDIR /src/app

RUN go install github.com/cosmtrek/air@v1.40.4

COPY . .

RUN go mod tidy