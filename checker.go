package goahttpcheck

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/ivpusic/httpcheck"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

type (
	decoder      = func(*http.Request) goahttp.Decoder
	encoder      = func(context.Context, http.ResponseWriter) goahttp.Encoder
	errorHandler = func(context.Context, http.ResponseWriter, error)
	middleware   = func(http.Handler) http.Handler

	HandlerBuilder func(goa.Endpoint, goahttp.Muxer, decoder, encoder, errorHandler) http.Handler
	HandlerMounter func(goahttp.Muxer, http.Handler)
)

type APIChecker struct {
	Mux          goahttp.Muxer
	Middleware   []middleware
	Decoder      decoder
	Encoder      encoder
	ErrorHandler errorHandler
}

type Option func(c *APIChecker)

func Muxer(mux goahttp.Muxer) Option {
	return func(c *APIChecker) {
		c.Mux = mux
	}
}

func Decoder(dec decoder) Option {
	return func(c *APIChecker) {
		c.Decoder = dec
	}
}

func Encoder(enc encoder) Option {
	return func(c *APIChecker) {
		c.Encoder = enc
	}
}

func ErrorHandler(eh errorHandler) Option {
	return func(c *APIChecker) {
		c.ErrorHandler = eh
	}
}

func New(options ...Option) *APIChecker {
	ret := &APIChecker{
		Mux:        goahttp.NewMuxer(),
		Middleware: []middleware{},
		Decoder:    goahttp.RequestDecoder,
		Encoder:    goahttp.ResponseEncoder,
		ErrorHandler: func(ctx context.Context, w http.ResponseWriter, err error) {
			log.Printf("ERROR: %v", err)
		},
	}
	for _, v := range options {
		v(ret)
	}
	return ret
}

func (c *APIChecker) Mount(builder HandlerBuilder, mounter HandlerMounter, endpoint goa.Endpoint) {
	handler := builder(endpoint, c.Mux, c.Decoder, c.Encoder, c.ErrorHandler)
	mounter(c.Mux, handler)
}
func (c *APIChecker) Use(middleware func(http.Handler) http.Handler) {
	c.Middleware = append(c.Middleware, middleware)
}

func (c APIChecker) Test(t *testing.T, method, path string) *httpcheck.Checker {
	var handler http.Handler = c.Mux
	for _, v := range c.Middleware {
		handler = v(handler)
	}
	checker := httpcheck.New(t, handler)
	return checker.Test(method, path)
}
