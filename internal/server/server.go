package server

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/repository"
)

type server struct {
	repository repository.Repository
	api.UnimplementedPaymentsServiceServer
}

func NewServer(repository repository.Repository) *server {
	return &server{
		repository: repository,
	}
}
