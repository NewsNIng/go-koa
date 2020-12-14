package god

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	Path string
	Method string

	Params map[string]string

	StatusCode int
	body []byte
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:    req,
		Writer: w,

		Path: req.URL.Path,
		Method: req.Method,

	}
}

func (ctx *Context) Body(body []byte) (n int, err error){
	ctx.body = body
	return 0, nil
}

func (ctx *Context) end() {
	ctx.Writer.Write(ctx.body)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}


func (ctx *Context) Write(b []byte)(n int, err error) {
	return ctx.Body(b)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}


func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]
	return value
}

func (ctx *Context) Error(message string) {
	ctx.Status(http.StatusInternalServerError)
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Writer.Write([]byte(message))
}

func (ctx *Context) JSON(obj interface{}) {
	ctx.Status(http.StatusOK)
	ctx.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.Status(code)
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}