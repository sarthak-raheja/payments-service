package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/processor"
	"github.com/sarthakraheja/payments-service/internal/repository"
	"github.com/sarthakraheja/payments-service/internal/server"
	"github.com/sarthakraheja/payments-service/internal/validator"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: Move these to a config file
const (
	port     = "9000"
	host     = "host.docker.internal"
	dbport   = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "paymentservice"
)

func main() {
	flag.Parse()
	logger := log.Default()
	logger.Printf("Listening on 9000")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, dbport, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatalf("Error setting database server %v:", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	validator := validator.NewValidator()
	processor := processor.NewProcessor(repository)

	server := server.NewServer(repository, validator, processor)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	api.RegisterPaymentsServiceServer(grpcServer, server)
	logger.Printf("Listening on 9000")
	grpcServer.Serve(lis)
}
