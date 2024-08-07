# Build stage
FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN go build -o main main.go

EXPOSE 12123

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 12123
CMD [ "/app/main" ]