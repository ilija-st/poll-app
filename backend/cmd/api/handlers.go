package main

import (
	"backend/ent"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
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

func (app *application) OneUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	user, err := app.DB.GetUserById(uid)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, user)
}

func (app *application) OnePoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	poll, err := app.DB.GetPollById(pid)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, poll)
}

func (app *application) CreatePoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var requestPayload struct {
		Question string   `json:"question"`
		Options  []string `json:"options"`
		UserId   int      `json:"user_id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Println(requestPayload)

	poll, err := app.DB.CreatePoll(requestPayload.Question, requestPayload.Options, requestPayload.UserId)
	if err != nil {
		app.errorJSON(w, errors.New("error when creating a poll"), http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, poll)
}

func (app *application) UpdatePoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var requestPayload struct {
		PollId  int      `json:"poll_id"`
		Options []string `json:"options"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Println("Request payload:", requestPayload)

	poll, err := app.DB.UpdatePoll(requestPayload.PollId, requestPayload.Options)
	if err != nil {
		app.errorJSON(w, errors.New("error when updating a poll"), http.StatusBadRequest)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, poll)
}

func (app *application) DeletePoll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var requestPayload struct {
		PollId int `json:"id"`
		UserId int `json:"user_id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate that the poll belongs to a user with given id
	poll, err := app.DB.GetPollById(requestPayload.PollId)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if poll.Edges.User.ID != requestPayload.UserId {
		app.errorJSON(w, errors.New("user with given id is not the owner of this poll"), http.StatusBadRequest)
		return
	}

	// delete a poll
	log.Println("Deleting poll:", poll)
	err = app.DB.DeletePoll(poll.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *application) VoteOnPollOption(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var requestPayload struct {
		UserId       int `json:"user_id"`
		PollId       int `json:"poll_id"`
		PollOptionId int `json:"poll_option_id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	log.Println("Request payload:", requestPayload)

	// 1. Validate poll option exists
	_, err = app.DB.GetPollOptionById(requestPayload.PollOptionId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// 2. Check if user has already voted on this poll
	poll, err := app.DB.GetPollById(requestPayload.PollId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Println("Poll:", poll)
	for _, pollOpt := range poll.Edges.PollOptions {
		votes, err := pollOpt.QueryVotes().All(context.Background())
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		for _, v := range votes {
			users, err := v.QueryUser().All(context.Background())
			if err != nil {
				app.writeJSON(w, http.StatusBadRequest, "Could not find users")
				return
			}
			for _, u := range users {
				if u.ID == requestPayload.UserId {
					log.Println("User already voted on this poll")
					app.writeJSON(w, http.StatusBadRequest, "User already voted on a poll.")
					return
				}
			}
		}

	}

	// 3. Fetch updated poll with vote counts
	app.DB.VoteOnPollOption(requestPayload.UserId, requestPayload.PollOptionId)

	poll, err = app.DB.GetPollById(requestPayload.PollId)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// 4. Respond with updated poll and vote information

	_ = app.writeJSON(w, http.StatusOK, poll)
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
	log.Println("Calling authenticate with: ", requestPayload)

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

	type Response struct {
		Token        string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		User         *ent.User `json:"user"`
	}

	res := Response{
		Token:        tokens.Token,
		RefreshToken: tokens.RefreshToken,
		User:         user,
	}

	log.Println(tokens.Token)
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	// this will send a cookie to whatever response we send
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, res)
}

func (app *application) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// read json payload
	var requestPayload struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	log.Println("Calling register with: ", requestPayload)

	// validate user against dataabse
	exists, err := app.DB.ExistsUserWithEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	if exists {
		app.errorJSON(w, errors.New("user already exists"), http.StatusBadRequest)
		return
	}

	// create a new user
	user, err := app.DB.CreateUser(requestPayload.FirstName, requestPayload.LastName, requestPayload.Email, requestPayload.Password)
	if err != nil {
		app.errorJSON(w, errors.New("error when creating a user"), http.StatusBadRequest)
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

	type Response struct {
		Token        string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		User         *ent.User `json:"user"`
	}

	res := Response{
		Token:        tokens.Token,
		RefreshToken: tokens.RefreshToken,
		User:         user,
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	// this will send a cookie to whatever response we send
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, res)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Calling refresh token function")
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			log.Println("Found jwt token cookie")
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserById(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			// TODO: Extract this part of code since its repeating in three functions
			type Response struct {
				Token        string    `json:"access_token"`
				RefreshToken string    `json:"refresh_token"`
				User         *ent.User `json:"user"`
			}

			res := Response{
				Token:        tokenPairs.Token,
				RefreshToken: tokenPairs.RefreshToken,
				User:         user,
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, res)
		}
	}
}
