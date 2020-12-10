package main

import (
	"fmt"
	"god"
)

func main() {
	app := god.New()

	app.Use(god.FilterFaviconIco())
	app.Use(god.Logger())




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
		// ctx.Body([]byte("good"))
		ctx.JSON(god.H{
			"name": "news",
			"age": 18,
		})
		fmt.Println("end 3")
	})

	app.Use(func(ctx *god.Context, next god.NextHandle) {
		fmt.Println("start 4")
		next()
		ctx.JSON(god.H{
			"name": "news4",
			"age": 184,
		})
		fmt.Println("end 4")
	})

	app.Listen(1998)
}
