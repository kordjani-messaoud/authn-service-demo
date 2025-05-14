# syntax=docker/dockerfile:1.7-labs

FROM golang:1.24-bookworm

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY --exclude=.gitignore --exclude=/bin --exclude=.env  ./backend /app/

RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]