package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"github.com/sarthakraheja/bank-simulator/server"

	"google.golang.org/grpc"
)

const (
	port = "9090"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("listening on 9090")

	acquiringBankServer := server.NewServer()
	grpcServer := grpc.NewServer()

	protos.RegisterAcquiringBankServiceServer(grpcServer, acquiringBankServer)

	grpcServer.Serve(lis)
}
