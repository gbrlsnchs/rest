package rest

import (
	"context"
	"net/http"
)

// ParamsFunc is a function to get URL parameter values.
type ParamsFunc func(context.Context) map[string]string

// Wrapper is a RESTful context wrapper.
type Wrapper struct {
	Handler        http.Handler
	RecoverHandler http.Handler
	ParamsFunc     ParamsFunc
}

func (wr *Wrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = &response{ResponseWriter: w, paramsFunc: wr.ParamsFunc}
	if wr.RecoverHandler != nil {
		defer wr.RecoverHandler.ServeHTTP(w, r)
	}
	wr.Handler.ServeHTTP(w, r)
}
