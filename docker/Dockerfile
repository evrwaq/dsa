FROM golang:1.24-alpine

WORKDIR /app

COPY ../cmd cmd
COPY ../internal internal
COPY ../go.mod go.mod
COPY ../go.sum go.sum

RUN go mod tidy

RUN go build -o main cmd/main.go

CMD ["./main"]
