package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"git.d.foundation/datcom/backend/src/domain"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (c *CoreHandler) GetGoogleLoginURL(w http.ResponseWriter, r *http.Request) {
	state := c.service.RandToken()
	url := googleOauthConfig.AuthCodeURL(state)

	json.NewEncoder(w).Encode(&domain.AuthConfig{
		RedirectURI: url,
		ClientID:    googleOauthConfig.ClientID,
	})
}

func (c *CoreHandler) VerifyGoogleUserLogin(w http.ResponseWriter, r *http.Request) {
	googleUser, err := c.service.GetUserDataFromGoogle(googleOauthConfig, r.FormValue("code"))
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	newToken, ok := r.URL.Query()["state"]
	if !ok || len(newToken[0]) < 1 {
		handleHTTPError(errors.New("Param 'state' not found"), http.StatusBadRequest, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}
	// Check user in db
	dbUser, _ := c.service.GetUserByEmail(tx, googleUser.Email)
	if dbUser != nil {
		if newToken[0] != dbUser.Token {
			// If logged out or login in new device, update new token to record
			err := c.service.UpdateUserToken(tx, &domain.CreateUserInput{
				Email: dbUser.Email,
				Token: dbUser.Token,
			}, newToken[0])

			if err != nil {
				tx.Rollback()
				handleHTTPError(err, http.StatusInternalServerError, w)
				return
			}
			dbUser, err = c.service.GetUserByEmail(tx, googleUser.Email)
			if err != nil {
				tx.Rollback()
				handleHTTPError(err, http.StatusInternalServerError, w)
				return
			}

		}
		tx.Commit()
		json.NewEncoder(w).Encode(domain.UserOutputMapping(dbUser))
		return
	}

	// if user is new, check with Fortress before decide to create user or not
	ftUser, err := c.service.FortressVerify(googleUser.Email)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	tx, err = c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	dbUser, err = c.service.CreateUser(tx, &domain.CreateUserInput{
		Name:  ftUser.Name,
		Email: ftUser.Email,
		Token: newToken[0],
	})
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	tx.Commit()
	json.NewEncoder(w).Encode(domain.UserOutputMapping(dbUser))
}
