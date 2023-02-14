package server

import (
	"context"
	"log"
	"math/rand"

	"github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct {
	protos.UnimplementedAcquiringBankServiceServer
}

func (s *server) CapturePayment(ctx context.Context, req *protos.CapturePaymentRequest) (*protos.CapturePaymentResponse, error) {
	logger := log.Default()
	logger.Printf("recieved request")

	rand := rand.Int() % 10
	if rand == 3 {
		return nil, grpc.Errorf(codes.Unknown, "unable to fulfil request")
	}

	return &protos.CapturePaymentResponse{}, nil
}

func (s *server) AuthorisePayment(ctx context.Context, req *protos.AuthorisePaymentRequest) (*protos.AuthorisePaymentResponse, error) {
	rand := rand.Int() % 10
	if rand == 1 || rand == 2 {
		return nil, grpc.Errorf(codes.Unknown, "unable to fulfil request")
	}

	return &protos.AuthorisePaymentResponse{}, nil
}

func NewServer() *server {
	return &server{}
}
