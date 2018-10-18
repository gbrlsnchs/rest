package rest_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/gbrlsnchs/rest"
)

func TestWrapper(t *testing.T) {
	testCases := []struct {
		wr     *Wrapper
		status int
	}{
		{
			wr: &Wrapper{
				Handler: HandlerFunc(func(ctx *Context) {
					if ctx.Param("hello") != "" {
						ctx.W.WriteHeader(http.StatusOK)
						return
					}
					ctx.W.WriteHeader(http.StatusBadRequest)
				}),
				ParamsFunc: func(_ context.Context) map[string]string {
					return map[string]string{"hello": "world"}
				},
			},
			status: http.StatusOK,
		},
		{
			wr: &Wrapper{
				Handler: HandlerFunc(func(ctx *Context) {
					if ctx.Param("hello") != "" {
						ctx.W.WriteHeader(http.StatusOK)
						return
					}
					ctx.W.WriteHeader(http.StatusBadRequest)
				}),
			},
			status: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			tc.wr.ServeHTTP(w, r)
			if want, got := tc.status, w.Code; want != got {
				t.Errorf("want %d, got %d", want, got)
			}
		})
	}
}
