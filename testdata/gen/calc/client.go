// Code generated by goa v3.21.1, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen calc/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	MultiplyEndpoint goa.Endpoint
	DivideEndpoint   goa.Endpoint
	RedirectEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(multiply, divide, redirect goa.Endpoint) *Client {
	return &Client{
		MultiplyEndpoint: multiply,
		DivideEndpoint:   divide,
		RedirectEndpoint: redirect,
	}
}

// Multiply calls the "multiply" endpoint of the "calc" service.
func (c *Client) Multiply(ctx context.Context, p *MultiplyPayload) (res int, err error) {
	var ires any
	ires, err = c.MultiplyEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Divide calls the "divide" endpoint of the "calc" service.
// Divide may return the following errors:
//   - "zero_division" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) Divide(ctx context.Context, p *DividePayload) (res int, err error) {
	var ires any
	ires, err = c.DivideEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Redirect calls the "redirect" endpoint of the "calc" service.
func (c *Client) Redirect(ctx context.Context) (err error) {
	_, err = c.RedirectEndpoint(ctx, nil)
	return
}
