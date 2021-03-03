package grpc

import (
	"github.com/westcottian/buyrent-breakevencalculator/domain"
	pb "github.com/westcottian/buyrent-breakevencalculator/proto/buyrent-breakevencalculator"
)

type CalculatorServiceServerImpl struct {
	calculatorServiceServer domain.BreakEvenService
}

func NewCalculatorServer(service domain.BreakEvenService) pb.BreakEvenServiceServer {
	return &CalculatorServiceServerImpl{
		calculatorServiceServer: service,
	}
}
