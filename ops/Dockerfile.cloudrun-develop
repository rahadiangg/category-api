FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build .


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/category-api .
CMD ["./category-api"]