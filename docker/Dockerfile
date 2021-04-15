FROM golang:1.15.1-alpine

COPY ./src /go/sim

WORKDIR /go/sim

RUN go build -o main main.go

EXPOSE 52200

RUN chmod 755 etc && chmod 666 etc/*

CMD ["./main"]