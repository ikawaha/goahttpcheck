package calcapi

import (
	"context"
	"errors"

	"calc/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct{}

// NewCalc returns the calc service implementation.
func NewCalc() calc.Service {
	return &calcsrvc{}
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
