// Code generated by goa v3.4.2, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen calc/design

package client

import (
	"fmt"
)

// AddCalcPath returns the URL path to the calc service add HTTP endpoint.
func AddCalcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}

// DivCalcPath returns the URL path to the calc service div HTTP endpoint.
func DivCalcPath(a int, b int) string {
	return fmt.Sprintf("/div/%v/%v", a, b)
}
