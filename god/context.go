package god

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter
	StatusCode int
	body []byte
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:    req,
		Writer: w,
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

func (ctx *Context) Next()  {

}

func (ctx *Context) Write(b []byte)(n int, err error) {
	return ctx.Body(b)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) JSON(obj interface{}) {
	ctx.Status(http.StatusOK)
	ctx.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Writer, err.Error(), 500)
	}
}