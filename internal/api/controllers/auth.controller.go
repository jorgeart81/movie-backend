package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

func (c *ApiController) RefreshToken(w http.ResponseWriter, r *http.Request) {

	if len(r.Cookies()) == 0 {
		c.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	for _, cookie := range r.Cookies() {
		if cookie.Name == c.Auth.CookieName {
			claims := &auth.Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
				return []byte(c.Auth.Secret), nil
			})
			if err != nil {
				c.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				c.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := c.Repository.GetUserByID(userID)
			if err != nil {
				c.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := auth.JwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := c.Auth.GenerateTokenPair(&u)
			if err != nil {
				c.errorJSON(w, errors.New("error generating token"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, c.Auth.GetRefreshCookie(tokenPairs.RefreshToken))

			c.writeJSON(w, http.StatusOK, tokenPairs)
		}
	}
}
