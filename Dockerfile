FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app ./cmd/order-service/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate ./cmd/migrate/main.go

FROM alpine AS webapp
WORKDIR /app
COPY --from=builder /go-app .
COPY --from=builder /app/config/config.yaml ./config/config.yaml
EXPOSE 50051
ENTRYPOINT ["./go-app"]

FROM alpine AS migrate
WORKDIR /app
COPY --from=builder /migrate .
COPY --from=builder /app/migrations ./migrations
ENTRYPOINT ["./migrate"]
