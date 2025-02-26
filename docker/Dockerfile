FROM golang:1.24-alpine

WORKDIR /app

COPY ../cmd cmd/
COPY ../internal internal/
COPY ../go.mod ./
# COPY ../go.sum go.sum TODO: Remove this comment later

RUN go mod tidy

RUN go build -o /app/main cmd/main.go

RUN chmod +x /app/main

CMD ["/app/main"]
