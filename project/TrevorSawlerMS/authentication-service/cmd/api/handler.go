package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJson(w, r, &payload)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}
	user, err := app.Models.User.GetByEmail(payload.Email)
	if err != nil {
		app.errorJson(w, errors.New("invalid login details"), http.StatusBadRequest)
		return
	}
	valid,err:=user.PasswordMatches(payload.Password)
	if err != nil || !valid{
		app.errorJson(w, errors.New("invalid login details"), http.StatusBadRequest)
		return
	}
	payloadData:=jsonResponse{
		Error: false,
		Message: fmt.Sprintf("User %s logged in.", user.Email),
		Data: user,
	}
	app.writeJson(w,http.StatusAccepted,payloadData)
}
