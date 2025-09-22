FROM golang:1.24.5-alpine

WORKDIR /app

COPY ./app /app

RUN go build -o personal-site

EXPOSE 8080

CMD ["./personal-site"]
