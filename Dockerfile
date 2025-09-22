FROM golang:1.21-alpine

WORKDIR /app

COPY ./app /app
COPY ./static /static

RUN go build -o personal-site

EXPOSE 8080

CMD ["./personal-site"]

