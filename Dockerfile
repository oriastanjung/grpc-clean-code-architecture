# Build Stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
WORKDIR /app/cmd
RUN go build -o main .

# Runtime Stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cmd/main .
EXPOSE 2701
CMD ["./main"]