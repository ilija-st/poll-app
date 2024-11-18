package main

import (
	"errors"
	"log"
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

func (app *application) authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against dataabse
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := PasswordMatches(requestPayload.Password, user.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// create a jwt user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(tokens.Token)
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	// this will send a cookie to whatever response we send
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, tokens)
}
