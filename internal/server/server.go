package server

import (
	"github.com/sarthakraheja/payments-service/api/v1/github.com/sarthakraheja/payments-service/api"
	"github.com/sarthakraheja/payments-service/internal/marshaller"
	"github.com/sarthakraheja/payments-service/internal/processor"
	"github.com/sarthakraheja/payments-service/internal/repository"
	"github.com/sarthakraheja/payments-service/internal/unmarshaller"
	"github.com/sarthakraheja/payments-service/internal/validator"
)

type server struct {
	repository   repository.Repository
	validator    validator.Validator
	processor    processor.Processor
	unmarshaller unmarshaller.Unmarshaller
	marshaller   marshaller.Marshaller
	api.UnimplementedPaymentsServiceServer
}

func NewServer(repository repository.Repository, validator validator.Validator, processor processor.Processor, unmarshaller unmarshaller.Unmarshaller, marshaller marshaller.Marshaller) *server {
	return &server{
		repository:   repository,
		validator:    validator,
		processor:    processor,
		unmarshaller: unmarshaller,
		marshaller:   marshaller,
	}
}
