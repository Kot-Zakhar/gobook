FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

ENV MONGO_CONNECTION_STRING 

RUN go build -o main ./cmd/gobook-server/main.go

CMD ["/app/main"]
