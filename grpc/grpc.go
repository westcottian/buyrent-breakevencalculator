package grpc

import (
	"github.com/westcottian/buyrent-breakevencalculator/domain"
	"google.golang.org/grpc"
)

// New creates a new gRPC server which handles coupon services.
func New(calculatorService domain.BreakEvenService) (*grpc.Server, error) {
	return newServer(calculatorService)
}

func newServer(calculatorService domain.BreakEvenService) (*grpc.Server, error) {
	// Create a new gRPC server.
	grpcServer := grpc.NewServer()

	// Initialize calculator server and register it to the gRPC server.
	calculatorServer := NewCalculatorServer(calculatorService)
	pb.RegisterBreakEvenServiceServer(calculatorServer)

	return grpcServer, nil
}
