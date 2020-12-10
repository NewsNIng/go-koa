package god

import (
	"log"
	"time"
)

func Logger() Middleware {
	return func(ctx *Context, next NextHandle) {
		// Start timer
		t := time.Now()
		next()
		log.Printf("[%d] [%s] %s in %v ", ctx.StatusCode, ctx.Req.Method, ctx.Req.RequestURI, time.Since(t))
	}
}

func FilterFaviconIco() Middleware {
	return func(ctx *Context, next NextHandle) {
		if ctx.Req.RequestURI == "/favicon.ico" {
			return
		} else {
			next()
		}
	}
}