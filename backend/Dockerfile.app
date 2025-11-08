FROM golang:1.25-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/osu-dashboard

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]