# Build stage
FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.cn
RUN go build -o main main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

EXPOSE 12123

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

RUN chmod +x /app/start.sh
RUN chmod +x /app/wait-for.sh

RUN ls -l /app

EXPOSE 12123
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]