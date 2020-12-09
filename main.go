package main

import (
	"fmt"
	"god"
)

func main() {
	app := god.New()

	filterFaviconIco := func(ctx *god.Context, next god.NextHandle) {
		if ctx.Req.RequestURI == "/favicon.ico" {
			return
		} else {
			next()
		}
	}
	app.Use(filterFaviconIco)

	app.Use(func(ctx *god.Context, next god.NextHandle) {
		fmt.Println("start 1")
		next()
		fmt.Println("end 1")
	})

	app.Use(func(ctx *god.Context, next god.NextHandle) {
		fmt.Println("start 2")
		next()
		fmt.Println("end 2")
	})

	app.Use(func(ctx *god.Context, next god.NextHandle) {
		fmt.Println("start 3")
		fmt.Println("ignore 4")
		// next()
		ctx.Body([]byte("good"))
		fmt.Println("end 3")
	})

	app.Use(func(ctx *god.Context, next god.NextHandle) {
		fmt.Println("start 4")
		next()
		fmt.Println("end 4")
	})

	app.Listen(1998)
}
