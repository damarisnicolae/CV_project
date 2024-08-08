FROM golang:1.22 AS start

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-api

FROM scratch

COPY --from=start /docker-api /docker-api

EXPOSE 8080

CMD [ "/docker-api" ]
