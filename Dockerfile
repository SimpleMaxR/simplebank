# Build stage
FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

EXPOSE 12123
CMD [ "/app/main" ]

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 12123
CMD [ "/app/main" ]