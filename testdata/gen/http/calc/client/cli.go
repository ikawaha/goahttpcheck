// Code generated by goa v3.16.0, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calc/design

package client

import (
	calc "calc/gen/calc"
	"fmt"
	"strconv"
)

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyA string, calcMultiplyB string) (*calc.MultiplyPayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcMultiplyA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcMultiplyB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.MultiplyPayload{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideA string, calcDivideB string) (*calc.DividePayload, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideA, 10, strconv.IntSize)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(calcDivideB, 10, strconv.IntSize)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &calc.DividePayload{}
	v.A = a
	v.B = b

	return v, nil
}
