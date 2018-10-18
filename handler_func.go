package rest

import "net/http"

// HandlerFunc is a function that holds a RESTful context.
type HandlerFunc func(ctx *Context)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	if rr, ok := w.(*response); ok && rr.paramsFunc != nil {
		ctx.SetParams(rr.paramsFunc)
	}
	h(ctx)
}
