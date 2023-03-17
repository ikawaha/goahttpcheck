// Code generated by goa v3.11.2, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen calc/design

package server

import (
	"fmt"
)

// MultiplyCalcPath returns the URL path to the calc service multiply HTTP endpoint.
func MultiplyCalcPath(a int, b int) string {
	return fmt.Sprintf("/multiply/%v/%v", a, b)
}

// DivideCalcPath returns the URL path to the calc service divide HTTP endpoint.
func DivideCalcPath(a int, b int) string {
	return fmt.Sprintf("/divide/%v/%v", a, b)
}

// RedirectCalcPath returns the URL path to the calc service redirect HTTP endpoint.
func RedirectCalcPath() string {
	return "/redirect"
}
