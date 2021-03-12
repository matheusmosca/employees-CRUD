FROM golang:1.15.10-alpine3.13

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go run ./out/api/main.go

# EXPOSE 8080

# CMD ["./out"]