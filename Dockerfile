FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o redis-cache

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/redis-cache .

CMD [ "./redis-cache" ]

