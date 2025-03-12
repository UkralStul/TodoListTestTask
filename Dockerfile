FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/todo-app .

COPY .env ./

EXPOSE 3000

CMD ["./todo-app"]