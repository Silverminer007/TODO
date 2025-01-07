FROM golang:1.23.4-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /app/app

CMD ["/app/app"]