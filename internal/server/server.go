package server

import "github.com/sarthakraheja/payments-service/internal/repository"

type server struct {
	repo repository.Repository
}

func NewServer(repo repository.Repository) *server {
	return &server{
		repo: repo,
	}
}
