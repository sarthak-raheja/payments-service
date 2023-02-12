package main

import (
	"fmt"
	"log"
	"net"

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

	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)

}
