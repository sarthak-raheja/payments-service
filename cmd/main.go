package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "9000"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	newServer := server.NewServer()
	api.RegisterPaymentsServiceServer(grpcServer, newServer)
	grpcServer.Serve(lis)

}
