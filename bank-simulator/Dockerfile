FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bank-simulator/bin/app bank-simulator/cmd/main.go

EXPOSE 9090

CMD ["./bank-simulator/bin/app"]
