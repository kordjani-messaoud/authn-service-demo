FROM golang:1.24-bookworm

WORKDIR /app

COPY ./src /app

RUN go mod download

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]