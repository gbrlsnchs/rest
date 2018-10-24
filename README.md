# rest (RESTful HTTP handler for Go)

[![Build Status](https://travis-ci.org/gbrlsnchs/rest.svg?branch=master)](https://travis-ci.org/gbrlsnchs/rest)
[![Sourcegraph](https://sourcegraph.com/github.com/gbrlsnchs/rest/-/badge.svg)](https://sourcegraph.com/github.com/gbrlsnchs/rest?badge)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/rest?status.svg)](https://godoc.org/github.com/gbrlsnchs/rest)
[![Minimal Version](https://img.shields.io/badge/minimal%20version-go1.10%2B-5272b4.svg)](https://golang.org/doc/go1.10)

## About
This package implements a simple RESTful HTTP handler that facilitates receiving requests or sending responses in JSON or XML by using a custom context.

## Usage
Full documentation [here](https://godoc.org/github.com/gbrlsnchs/rest).

### Installing
#### Go 1.10
`vgo get -u github.com/gbrlsnchs/rest`
#### Go 1.11 or after
`go get -u github.com/gbrlsnchs/rest`

### Importing
```go
import (
	// ...

	"github.com/gbrlsnchs/rest"
)
```

### Setting a wrapper
#### Consider the following type to be received / sent
```go
type message struct {
	content string `json:"content,omitempty"`
}
```

#### Now, set the main handler
```go
http.Handle("/", &rest.Wrapper{
	Handler: rest.HandlerFunc(func(ctx *rest.Context) {
		var ping message
		if err := ctx.ReceiveJSON(&ping); err != nil {
			// handle error
		}
		if ping.content != "ping" {
			ctx.Send(http.StatusBadRequest)
			return
		}
		pong := message{"pong"}
		ctx.SendJSON(pong, http.StatusOK)
	}),
	RecoverHandler: rest.HandlerFunc(func(ctx *rest.Context) {
		ctx.Send(http.StatusInternalServerError)
	}),
})
```

## Contributing
### How to help
- For bugs and opinions, please [open an issue](https://github.com/gbrlsnchs/rest/issues/new)
- For pushing changes, please [open a pull request](https://github.com/gbrlsnchs/rest/compare)
