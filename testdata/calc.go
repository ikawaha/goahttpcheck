package calcapi

import (
	"context"
	"errors"
	"log"

	"calc/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	return p.A * p.B, nil
}

// Divide implements divide.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res int, err error) {
	if p.B == 0 {
		return 0, calc.MakeZeroDivision(errors.New("zero division error"))
	}
	return p.A / p.B, nil
}

// Redirect implements redirect.
func (s *calcsrvc) Redirect(ctx context.Context) (err error) {
	return
}
