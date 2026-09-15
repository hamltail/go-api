FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/server ./server
COPY data ./data

CMD ["./server"]
