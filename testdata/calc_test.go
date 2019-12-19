package calcapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"calc/gen/calc"
	"calc/gen/http/calc/server"
	"github.com/ikawaha/goahttpcheck"
	goa "goa.design/goa/v3/pkg"
)

func TestCalcsrvc_Add(t *testing.T) {
	checker := goahttpcheck.New()
	var logger log.Logger
	checker.Mount(server.NewAddHandler, server.MountAddHandler, calc.NewAddEndpoint(NewCalc(&logger)))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, http.MethodGet, "/add/1/2").
		Check().
		HasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			r.Body.Close()
			if got, expected := strings.TrimSpace(string(b)), "3"; got != expected {
				t.Errorf("got %+v, expected %v", b, expected)
			}
		})
}

func TestCalcsrvc_Div(t *testing.T) {
	checker := goahttpcheck.New()

	var logger log.Logger
	checker.Mount(server.NewDivHandler, server.MountDivHandler, calc.NewDivEndpoint(NewCalc(&logger)))

	// see. https://github.com/ikawaha/httpcheck
	checker.Test(t, "GET", "/div/1/0").
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
