package main

import (
	"god"
	"net/http"
)

func main() {
	app := god.Default()
	app.Use(initRoutes())
	app.Listen(1998)
}

func initRoutes() god.Middleware {
	r := god.NewRouter()
	r.GET("/test/:name", func(ctx *god.Context) {
		//name := ctx.Param("name")
		//age := ctx.Query("age")
		names := []string{"test error"}
		ctx.String(http.StatusOK, names[100])
	}).GET("/test2", func(ctx *god.Context) {
		ctx.JSON(god.H{
			"name": "test2",
		})
	})
	return r.GetRoutes()
}
