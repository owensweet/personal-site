FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./app /app
COPY ./static /static

RUN go build -o personal-site ./app/main.go

EXPOSE 8080

CMD ["./personal-site"]

