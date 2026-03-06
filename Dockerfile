FROM golang:1.26-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o app ./cmd/server

EXPOSE 8080

# ENV GIN_MODE=release

CMD ["./app"]