FROM golang:1.23

WORKDIR /app

COPY ./app/go.mod ./

RUN go mod download

COPY ./app/ ./
COPY ./static ./static

# DEBUG
RUN echo "=== Files in /app ===" && ls -la /app
RUN echo "=== Files in /app/static ===" && ls -la /app/static || echo "static dir not found"
RUN echo "=== Looking for index.html ===" && find /app -name "index.html" || echo "index.html not found"

RUN go build -o personal-site

EXPOSE 8080

CMD ["./personal-site"]
