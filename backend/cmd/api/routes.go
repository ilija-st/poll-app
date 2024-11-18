package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.GET("/", app.Home)

	mux.GET("/polls", app.AllPolls)

	mux.GET("/users", app.AllUsers)

	mux.POST("/authenticate", app.authenticate)

	mux.GET("/refresh", app.refreshToken)

	mux.GET("/users/:id", app.OneUser)

	handler := app.enableCORS(mux)
	handler = app.Logger(handler)

	return handler
}
