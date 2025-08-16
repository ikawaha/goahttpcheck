package calcapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"calc/gen/calc"
	"calc/gen/http/calc/server"
	"github.com/ikawaha/goahttpcheck"
	goa "goa.design/goa/v3/pkg"
)

func TestCalcsrvc_Multiply(t *testing.T) {
	checker := goahttpcheck.New()
	checker.Mount(server.NewMultiplyHandler, server.MountMultiplyHandler, calc.NewMultiplyEndpoint(NewCalc()))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, http.MethodGet, "/multiply/1/2").
		Check().
		HasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			r.Body.Close()
			if got, expected := strings.TrimSpace(string(b)), "2"; got != expected {
				t.Errorf("got %+v, expected %v", b, expected)
			}
		})
}

func TestCalcsrvc_Divide(t *testing.T) {
	checker := goahttpcheck.New()

	checker.Mount(server.NewDivideHandler, server.MountDivideHandler, calc.NewDivideEndpoint(NewCalc()))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, "GET", "/divide/1/0").
		WithHeader("Host", "test.example.jp").
		WithHeader("Test", "test.example.jp").
		Check().
		HasStatus(http.StatusBadRequest).
		Cb(func(r *http.Response) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			r.Body.Close()
			var resp goa.ServiceError
			if err := json.Unmarshal(b, &resp); err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			if got, expected := resp.Message, "zero division error"; got != expected {
				t.Errorf("got %+v, expected %v", got, expected)
			}
		})
}

func TestCalcsrvc_Redirect(t *testing.T) {
	checker := goahttpcheck.New(goahttpcheck.NoRedirect())

	checker.Mount(server.NewRedirectHandler, server.MountRedirectHandler, calc.NewRedirectEndpoint(NewCalc()))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, "GET", "/redirect").
		Check().
		HasStatus(http.StatusTemporaryRedirect)
}
