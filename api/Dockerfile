FROM golang:1.22-alpine AS builder

ENV DB_HOST  db
ENV DB_PORT  3306

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.20.2

RUN apk --no-cache add wget \
    && wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz \
    && rm dockerize-linux-amd64-v0.6.1.tar.gz \
    && apk del wget

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["dockerize", "-wait", "tcp://cv_db-container:3306", "-timeout", "60s", "./main"]
