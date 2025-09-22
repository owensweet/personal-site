FROM golang:1.24.5

WORKDIR /app

COPY ./app /app
COPY ./static /static

RUN go build -o personal-site

EXPOSE 8080

CMD ["./personal-site"]

