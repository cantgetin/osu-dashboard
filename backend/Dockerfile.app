FROM golang:1.23-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/playcount-monitor

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]