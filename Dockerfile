FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN go build -o main .

RUN if [ ! -f "certs/cert.pem" ] || [ ! -f "certs/key.pem" ]; then \
        echo "Certificates not found. Generating..."; \
        openssl req -x509 -newkey rsa:4096 -keyout "certs/key.pem" -out "certs/cert.pem" -days 365 -nodes -subj '/CN=localhost'; \
    fi

RUN chmod 600 certs/cert.pem certs/key.pem

EXPOSE 1321

CMD ["./main"]