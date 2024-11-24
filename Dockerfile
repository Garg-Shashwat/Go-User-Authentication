FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service main.go

EXPOSE 9000

CMD ["./auth-service"]