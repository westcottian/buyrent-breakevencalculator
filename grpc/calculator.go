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

// ComputePropertyBreakEven calls the core service's ComputePropertyBreakEven method and maps the result to a grpc service response.
func (c *CalculatorServiceServerImpl) CalculatePropertyBreakEven(ctx context.Context, req *pb.BreakEvenRequest) (*pb.BreakEvenResponse, error) {
	//Get the request params and validate them.
	params := domain.BuyRentCalculator{}

	err := c.calculatorServiceServer.CalculatePropertyBreakEven(params)
	if err != nil {
		return nil, err
	}

	/*result = &appgrpc.BreakEvenResponse{
		Rent:     res.Rent,
		Purchase: res.Purchase,
		Message:  res.Message,
		Verdict:  res.Verdict,
	}*/

	return nil, nil
}
