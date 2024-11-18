package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Poll backend up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllPolls(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	polls, err := app.DB.AllPolls()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, polls)
}

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := app.DB.AllUsers()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, users)
}
