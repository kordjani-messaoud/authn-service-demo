# syntax=docker/dockerfile:1.7-labs

FROM golang:1.24-bookworm

WORKDIR /app

COPY --exclude=.gitignore --exclude=/bin --exclude=.env  . /app

RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]