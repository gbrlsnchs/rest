package rest

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Context is a RESTful context.
type Context struct {
	W    http.ResponseWriter
	R    *http.Request
	pmap map[string]string // params map
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:    w,
		R:    r,
		pmap: make(map[string]string),
	}
}

// Param returns the parameter value, if any.
func (c *Context) Param(key string) string {
	return c.pmap[key]
}

// ReceiveJSON reads a JSON request and sets it to an address.
func (c *Context) ReceiveJSON(v interface{}) error {
	return c.receive(json.Unmarshal, v)
}

// ReceiveXML reads a XML request and sets it to an address.
func (c *Context) ReceiveXML(v interface{}) error {
	return c.receive(xml.Unmarshal, v)
}

// Send sends a bodiless response.
func (c *Context) Send(status int) {
	c.send(nil, status)
}

// SendJSON sends a response encoded to JSON.
func (c *Context) SendJSON(body interface{}, status int) {
	c.W.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(body)
	c.send(b, status)
}

// SendXML sends a response encoded to XML.
func (c *Context) SendXML(body interface{}, status int) {
	c.W.Header().Set("Content-Type", "application/xml")
	b, _ := xml.Marshal(body)
	c.send(b, status)
}

// SetParams sets the URL parameter values based on a function.
func (c *Context) SetParams(fn ParamsFunc) {
	if fn != nil {
		c.pmap = fn(c.R.Context())
	}
}

func (c *Context) send(body []byte, status int) {
	c.W.WriteHeader(status)
	c.W.Write(body)
}

func (c *Context) receive(unmarshal func([]byte, interface{}) error, v interface{}) error {
	b, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}
	return unmarshal(b, v)
}
