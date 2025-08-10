# FROM golang:1.24.6-alpine3.22
# FROM golang:1.25rc3-bullseye
# FROM golang:1.25-rc-bookworm
FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
# RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN goose up

# RUN swag init
RUN go build -o ./tmp/main ./cmd/main.go
# CMD ["./tmp/main"]

