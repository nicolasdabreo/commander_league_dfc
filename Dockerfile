FROM golang:1.21.6-bullseye

ENV GIN_MODE=release

WORKDIR /app

# Install Go dependencies
COPY go.mod go.sum ./
RUN go mod download


# Install SQLite dependencies
RUN apt-get update && \
    apt-get install -y sqlite3

# Copy codebase
COPY . .

# Build go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

EXPOSE 8080

CMD ["./main"]