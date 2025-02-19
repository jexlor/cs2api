FROM golang:1.24.0-alpine

RUN apk add --no-cache git bash && \
    go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air"]
