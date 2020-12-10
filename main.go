package main

import (
	"god"
)

func main() {
	app := god.New()

	app.Use(god.FilterFaviconIco())
	app.Use(god.Logger())
	app.Use(initRoutes())

	app.Listen(1998)
}

func initRoutes() god.Middleware {
	r := god.NewRouter()
	r.GET("/test/:name", func(ctx *god.Context) {
		name := ctx.Param("name")
		age := ctx.Query("age")
		ctx.JSON(god.H{
			"name": name,
			"age":  age,
		})
	}).GET("/test2", func(ctx *god.Context) {
		ctx.JSON(god.H{
			"name": "test2",
		})
	})
	return r.GetRoutes()
}
