package controllers

import (
	"errors"
	"net/http"

	"github.com/jorgeart81/movie-backend/internal/api/auth"
)

func (c *ApiController) Authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := c.Repository.GetUserByEmail(requestPayload.Email)
	if err != nil {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// check password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	// create jwt user
	u := auth.JwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// generate tokens
	tokens, err := c.Auth.GenerateTokenPair(&u)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	refreshCookie := c.Auth.GetRefreshCookie(tokens.Token)
	http.SetCookie(w, refreshCookie)

	c.writeJSON(w, http.StatusOK, tokens)
}
