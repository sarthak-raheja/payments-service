package server

import "github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"

type server struct {
	api.UnimplementedPaymentsServiceServer
}

func NewServer() *server {
	return &server{}
}
