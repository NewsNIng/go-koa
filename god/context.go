package god

import "net/http"

type Context struct {
	Req    *http.Request
	Writer http.ResponseWriter

	body []byte
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Req:    req,
		Writer: w,
	}
}

func (ctx *Context) Body(body []byte) {
	ctx.body = body
}

func (ctx *Context) end() {
	ctx.Writer.Write(ctx.body)
}
