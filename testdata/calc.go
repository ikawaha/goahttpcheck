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

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	return p.A + p.B, nil
}

// Div implements div.
func (s *calcsrvc) Div(ctx context.Context, p *calc.DivPayload) (res int, err error) {
	if p.B == 0 {
		return 0, calc.MakeZeroDivision(errors.New("zero division error"))
	}
	return p.A / p.B, nil
}

// Redirect implements redirect.
func (s *calcsrvc) Redirect(ctx context.Context) (err error) {
	return
}
