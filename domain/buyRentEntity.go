package domain

type BuyRentCalculator struct {
	PropertyID            int64
	TotalPropertyCost     int64
	StayTerm              int64
	DownPayment           float64
	InterestRate          float64
	PropertyTaxes         float64
	PropertyTransferTaxes float64
	LoanTerm              int64
	MonthlyRent           int64
}

type BreakEvenService interface {
	CalculatePropertyBreakEven(req BuyRentCalculator) (err error)
}
