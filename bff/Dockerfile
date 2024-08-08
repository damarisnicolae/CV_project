FROM golang:1.22 

WORKDIR /bff

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /bff-app

EXPOSE 8080

CMD ["/bff-app"]
