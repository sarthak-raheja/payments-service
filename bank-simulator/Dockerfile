FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/app cmd/main.go

EXPOSE 9090

CMD ["./bin/app"]
