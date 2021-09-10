// Code generated by goa v3.5.1, DO NOT EDIT.
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
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "calc" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.AddPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "add", "*calc.AddPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeAddResponse returns a decoder for responses returned by the calc add
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
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
				return nil, goahttp.ErrDecodingError("calc", "add", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "add", resp.StatusCode, string(body))
		}
	}
}

// BuildDivRequest instantiates a HTTP request object with method and path set
// to call the "calc" service "div" endpoint
func (c *Client) BuildDivRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.DivPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "div", "*calc.DivPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DivCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "div", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDivResponse returns a decoder for responses returned by the calc div
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
// DecodeDivResponse may return the following errors:
//	- "zero_division" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeDivResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
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
				return nil, goahttp.ErrDecodingError("calc", "div", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body DivZeroDivisionResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "div", err)
			}
			err = ValidateDivZeroDivisionResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "div", err)
			}
			return nil, NewDivZeroDivision(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "div", resp.StatusCode, string(body))
		}
	}
}

// BuildRedirectRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "redirect" endpoint
func (c *Client) BuildRedirectRequest(ctx context.Context, v interface{}) (*http.Request, error) {
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
func DecodeRedirectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusTemporaryRedirect:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "redirect", resp.StatusCode, string(body))
		}
	}
}
