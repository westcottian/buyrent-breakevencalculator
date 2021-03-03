package grpc

import (
	"github.com/westcottian/buyrent-breakevencalculator/domain"
	pbbr "github.com/westcottian/buyrent-breakevencalculator/pb"
)

type CalculatorServiceServerImpl struct {
	calculatorServiceServer domain.BreakEvenService
}

func NewCalculatorServer(service domain.BreakEvenService) pbbr.BreakEvenServiceServer {
	return &CalculatorServiceServerImpl{
		calculatorServiceServer: service,
	}
}

// ComputePropertyBreakEven calls the core service's ComputePropertyBreakEven method and maps the result to a grpc service response.
func (c *CalculatorServiceServerImpl) CalculatePropertyBreakEven(ctx context.Context, req *pbbr.BreakEvenRequest) (*pbbr.BreakEvenResponse, error) {
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
