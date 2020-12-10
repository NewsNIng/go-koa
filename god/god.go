package god

import (
	"net/http"
	"strconv"
)

type NextHandle func()
type Middleware func(ctx *Context, next NextHandle)
type H map[string]interface{}

type Application struct {
	middlewares []Middleware
}

func New() *Application {
	return &Application{}
}

func (app *Application) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	app.dispatchMiddleware(ctx, 0)
	ctx.end()
}

func (app *Application) Listen(addr int) error {
	return http.ListenAndServe(":"+strconv.Itoa(addr), app)
}

func (app *Application) Use(middleware Middleware) {
	app.middlewares = append(app.middlewares, middleware)
}

func (app *Application) dispatchMiddleware(ctx *Context, index int) {
	if index == len(app.middlewares) {
		return
	}
	currentMiddleware := app.middlewares[index]
	currentMiddleware(ctx, func() {
		app.dispatchMiddleware(ctx, index+1)
	})
}
