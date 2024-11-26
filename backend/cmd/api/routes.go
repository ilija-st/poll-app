package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.GET("/", app.Home)

	mux.GET("/polls", app.AllPolls)

	mux.POST("/polls", app.authRequired(app.CreatePoll))

	mux.PUT("/polls", app.authRequired(app.UpdatePoll))

	mux.DELETE("/polls", app.authRequired(app.DeletePoll))

	mux.GET("/polls/:id", app.OnePoll)

	mux.PUT("/polls/:id", app.authRequired(app.VoteOnPollOption))

	mux.GET("/users", app.authRequired(app.AllUsers))

	mux.POST("/authenticate", app.authenticate)

	mux.POST("/register", app.register)

	mux.GET("/logout", app.logout)

	mux.GET("/refresh", app.refreshToken)

	mux.GET("/users/:id", app.authRequired(app.OneUser))

	handler := app.enableCORS(mux)
	handler = app.Logger(handler)

	return handler
}
