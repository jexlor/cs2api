FROM golang:1.22.0-alpine 
#also change this config to 1.23-alpine
RUN apk add --no-cache bash curl
WORKDIR /app
# RUN go install github.com/air-verse/air@latest update to go 1.23
COPY go.mod go.sum ./
RUN go mod download
COPY . .
EXPOSE 8080
CMD ["air"]
