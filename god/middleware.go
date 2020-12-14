package god

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func Logger() Middleware {
	return func(ctx *Context, next NextHandle) {
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



// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(4, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}


func Recovery() Middleware {
	return func(ctx *Context, next NextHandle) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				ctx.Error("Internal Server Error")
			}
		}()
		next()
	}
}