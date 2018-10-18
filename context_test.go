package rest_test

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/gbrlsnchs/rest"
)

type responseMock struct {
	XMLName xml.Name `json:"-" xml:"mock"`
	Hello   string   `json:"hello,omitempty" xml:"hello"`
}

func TestReceiveJSON(t *testing.T) {
	var mock responseMock
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(`{"hello":"world"}`))
	HandlerFunc(func(ctx *Context) {
		ctx.ReceiveJSON(&mock)
	}).ServeHTTP(w, r)
	if want, got := "world", mock.Hello; want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestReceiveXML(t *testing.T) {
	var mock responseMock
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("<mock><hello>world</hello></mock>"))
	HandlerFunc(func(ctx *Context) {
		ctx.ReceiveXML(&mock)
	}).ServeHTTP(w, r)
	if want, got := "world", mock.Hello; want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestContextSend(t *testing.T) {
	testCases := []struct {
		status int
	}{
		{http.StatusAccepted},
		{http.StatusAlreadyReported},
		{http.StatusBadGateway},
		{http.StatusBadRequest},
		{http.StatusConflict},
		{http.StatusContinue},
		{http.StatusCreated},
		{http.StatusExpectationFailed},
		{http.StatusFailedDependency},
		{http.StatusForbidden},
		{http.StatusFound},
		{http.StatusGatewayTimeout},
		{http.StatusGone},
		{http.StatusHTTPVersionNotSupported},
		{http.StatusIMUsed},
		{http.StatusInsufficientStorage},
		{http.StatusInternalServerError},
		{http.StatusLengthRequired},
		{http.StatusLocked},
		{http.StatusLoopDetected},
		{http.StatusMethodNotAllowed},
		{http.StatusMovedPermanently},
		{http.StatusMultiStatus},
		{http.StatusMultipleChoices},
		{http.StatusNetworkAuthenticationRequired},
		{http.StatusNoContent},
		{http.StatusNonAuthoritativeInfo},
		{http.StatusNotAcceptable},
		{http.StatusNotExtended},
		{http.StatusNotExtended},
		{http.StatusNotFound},
		{http.StatusNotImplemented},
		{http.StatusNotModified},
		{http.StatusOK},
		{http.StatusPartialContent},
		{http.StatusPaymentRequired},
		{http.StatusPermanentRedirect},
		{http.StatusPreconditionFailed},
		{http.StatusPreconditionRequired},
		{http.StatusProcessing},
		{http.StatusProxyAuthRequired},
		{http.StatusRequestEntityTooLarge},
		{http.StatusRequestHeaderFieldsTooLarge},
		{http.StatusRequestTimeout},
		{http.StatusRequestURITooLong},
		{http.StatusRequestedRangeNotSatisfiable},
		{http.StatusResetContent},
		{http.StatusSeeOther},
		{http.StatusServiceUnavailable},
		{http.StatusSwitchingProtocols},
		{http.StatusTeapot},
		{http.StatusTemporaryRedirect},
		{http.StatusTooManyRequests},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s (%d)", http.StatusText(tc.status), tc.status), func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			HandlerFunc(func(ctx *Context) {
				ctx.Send(tc.status)
			}).ServeHTTP(w, r)
			if want, got := tc.status, w.Code; want != got {
				t.Errorf("want %d, got %d", want, got)
			}
			if want, got := "", w.Body.Bytes(); want != string(got) {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestContextSendJSON(t *testing.T) {
	v := responseMock{
		Hello: "world",
	}
	vj := `{"hello":"world"}`
	testCases := []struct {
		status int
	}{
		{http.StatusAccepted},
		{http.StatusAlreadyReported},
		{http.StatusBadGateway},
		{http.StatusBadRequest},
		{http.StatusConflict},
		{http.StatusContinue},
		{http.StatusCreated},
		{http.StatusExpectationFailed},
		{http.StatusFailedDependency},
		{http.StatusForbidden},
		{http.StatusFound},
		{http.StatusGatewayTimeout},
		{http.StatusGone},
		{http.StatusHTTPVersionNotSupported},
		{http.StatusIMUsed},
		{http.StatusInsufficientStorage},
		{http.StatusInternalServerError},
		{http.StatusLengthRequired},
		{http.StatusLocked},
		{http.StatusLoopDetected},
		{http.StatusMethodNotAllowed},
		{http.StatusMovedPermanently},
		{http.StatusMultiStatus},
		{http.StatusMultipleChoices},
		{http.StatusNetworkAuthenticationRequired},
		{http.StatusNoContent},
		{http.StatusNonAuthoritativeInfo},
		{http.StatusNotAcceptable},
		{http.StatusNotExtended},
		{http.StatusNotExtended},
		{http.StatusNotFound},
		{http.StatusNotImplemented},
		{http.StatusNotModified},
		{http.StatusOK},
		{http.StatusPartialContent},
		{http.StatusPaymentRequired},
		{http.StatusPermanentRedirect},
		{http.StatusPreconditionFailed},
		{http.StatusPreconditionRequired},
		{http.StatusProcessing},
		{http.StatusProxyAuthRequired},
		{http.StatusRequestEntityTooLarge},
		{http.StatusRequestHeaderFieldsTooLarge},
		{http.StatusRequestTimeout},
		{http.StatusRequestURITooLong},
		{http.StatusRequestedRangeNotSatisfiable},
		{http.StatusResetContent},
		{http.StatusSeeOther},
		{http.StatusServiceUnavailable},
		{http.StatusSwitchingProtocols},
		{http.StatusTeapot},
		{http.StatusTemporaryRedirect},
		{http.StatusTooManyRequests},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s (%d)", http.StatusText(tc.status), tc.status), func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			HandlerFunc(func(ctx *Context) {
				ctx.SendJSON(v, tc.status)
			}).ServeHTTP(w, r)
			if want, got := tc.status, w.Code; want != got {
				t.Errorf("want %d, got %d", want, got)
			}
			if want, got := "application/json", w.HeaderMap.Get("Content-Type"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
			if want, got := vj, w.Body.Bytes(); want != string(got) {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestContextSendXML(t *testing.T) {
	v := responseMock{
		Hello: "world",
	}
	vx := "<mock><hello>world</hello></mock>"
	testCases := []struct {
		status int
	}{
		{http.StatusAccepted},
		{http.StatusAlreadyReported},
		{http.StatusBadGateway},
		{http.StatusBadRequest},
		{http.StatusConflict},
		{http.StatusContinue},
		{http.StatusCreated},
		{http.StatusExpectationFailed},
		{http.StatusFailedDependency},
		{http.StatusForbidden},
		{http.StatusFound},
		{http.StatusGatewayTimeout},
		{http.StatusGone},
		{http.StatusHTTPVersionNotSupported},
		{http.StatusIMUsed},
		{http.StatusInsufficientStorage},
		{http.StatusInternalServerError},
		{http.StatusLengthRequired},
		{http.StatusLocked},
		{http.StatusLoopDetected},
		{http.StatusMethodNotAllowed},
		{http.StatusMovedPermanently},
		{http.StatusMultiStatus},
		{http.StatusMultipleChoices},
		{http.StatusNetworkAuthenticationRequired},
		{http.StatusNoContent},
		{http.StatusNonAuthoritativeInfo},
		{http.StatusNotAcceptable},
		{http.StatusNotExtended},
		{http.StatusNotExtended},
		{http.StatusNotFound},
		{http.StatusNotImplemented},
		{http.StatusNotModified},
		{http.StatusOK},
		{http.StatusPartialContent},
		{http.StatusPaymentRequired},
		{http.StatusPermanentRedirect},
		{http.StatusPreconditionFailed},
		{http.StatusPreconditionRequired},
		{http.StatusProcessing},
		{http.StatusProxyAuthRequired},
		{http.StatusRequestEntityTooLarge},
		{http.StatusRequestHeaderFieldsTooLarge},
		{http.StatusRequestTimeout},
		{http.StatusRequestURITooLong},
		{http.StatusRequestedRangeNotSatisfiable},
		{http.StatusResetContent},
		{http.StatusSeeOther},
		{http.StatusServiceUnavailable},
		{http.StatusSwitchingProtocols},
		{http.StatusTeapot},
		{http.StatusTemporaryRedirect},
		{http.StatusTooManyRequests},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s (%d)", http.StatusText(tc.status), tc.status), func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			HandlerFunc(func(ctx *Context) {
				ctx.SendXML(v, tc.status)
			}).ServeHTTP(w, r)
			if want, got := tc.status, w.Code; want != got {
				t.Errorf("want %d, got %d", want, got)
			}
			if want, got := "application/xml", w.HeaderMap.Get("Content-Type"); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
			if want, got := vx, w.Body.Bytes(); want != string(got) {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
