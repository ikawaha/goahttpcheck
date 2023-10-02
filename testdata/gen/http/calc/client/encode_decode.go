// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen calc/design

package client

import (
	"bytes"
	calc "calc/gen/calc"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildMultiplyRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "multiply" endpoint
func (c *Client) BuildMultiplyRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.MultiplyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiplyCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "multiply", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeMultiplyResponse returns a decoder for responses returned by the calc
// multiply endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeMultiplyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "multiply", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "multiply", resp.StatusCode, string(body))
		}
	}
}

// BuildDivideRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "divide" endpoint
func (c *Client) BuildDivideRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.DividePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "divide", "*calc.DividePayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DivideCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "divide", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDivideResponse returns a decoder for responses returned by the calc
// divide endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDivideResponse may return the following errors:
//   - "zero_division" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeDivideResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body DivideZeroDivisionResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			err = ValidateDivideZeroDivisionResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "divide", err)
			}
			return nil, NewDivideZeroDivision(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "divide", resp.StatusCode, string(body))
		}
	}
}

// BuildRedirectRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "redirect" endpoint
func (c *Client) BuildRedirectRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RedirectCalcPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "redirect", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRedirectResponse returns a decoder for responses returned by the calc
// redirect endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRedirectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusTemporaryRedirect:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "redirect", resp.StatusCode, string(body))
		}
	}
}
