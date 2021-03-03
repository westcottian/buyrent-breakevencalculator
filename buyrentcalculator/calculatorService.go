package buyrentcalculator

import (
	"github.com/westcottian/buyrent-breakevencalculator/domain"
)

type service struct {
	data domain.BuyRentCalculator
}

// NewService instantiates a new Service.
func NewService() domain.BreakEvenService {
	return &service{
		data: domain.BuyRentCalculator{},
	}
}
