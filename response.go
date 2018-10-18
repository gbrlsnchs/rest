package rest

import (
	"net/http"
)

type response struct {
	http.ResponseWriter
	paramsFunc ParamsFunc
}
