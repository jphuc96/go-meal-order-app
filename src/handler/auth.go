package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/src/domain"
)

const (
	oauthGoogleUrlAPI = "https://wwg.Writer.googleapis.com/oauth2/v2/userinfo?access_token="
)

func (c *CoreHandler) GoogleLogin(g *gin.Context) {
	State = c.service.RandToken()
	url := OAuthConfig.AuthCodeURL(State)

	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&domain.AuthConfig{
		RedirectURL: url,
		ClientID:    OAuthConfig.ClientID,
		State:       State,
	})
}

func (c *CoreHandler) GoogleLogout(g *gin.Context) {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	err = c.service.UpdateUserToken(tx, &domain.UserInput{
		Email: g.Request.Header.Get("email"),
		Token: g.Request.Header.Get("access_token"),
	}, "")
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	tx.Commit()
	g.Writer.WriteHeader(http.StatusOK)
}

func (c *CoreHandler) GoogleOauthCallback(g *gin.Context) {
	queryState := g.Request.Header.Get("state")
	if State != queryState {
		c.HandleHTTPError(domain.InvalidOAuthState, http.StatusUnauthorized, g.Writer)
		return
	}

	email := g.Request.Header.Get("email")
	if email == "" {
		c.HandleHTTPError(domain.NotProvideEmail, http.StatusUnauthorized, g.Writer)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	// Check user in db
	dbUser, _ := c.service.GetUserByEmail(tx, email)
	if dbUser != nil {
		// if logged out
		if dbUser.Token == "" {
			err := c.service.UpdateUserToken(tx, &domain.UserInput{
				Email: dbUser.Email,
				Token: dbUser.Token,
			}, c.service.RandToken())
			if err != nil {
				tx.Rollback()
				c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
				return
			}
			dbUser, _ = c.service.GetUserByEmail(tx, email)
		}

		tx.Commit()
		json.NewEncoder(g.Writer).Encode(domain.UserOutputMapping(dbUser))
		return
	}

	// if user is new, check with Fortress before decide to create user or not
	ftUser, err := c.service.FortressVerify(email)
	if err != nil {
		c.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}

	tx, err = c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	dbUser, err = c.service.CreateUser(tx, &domain.UserInput{
		Name:  ftUser.Name,
		Email: ftUser.Email,
		Token: c.service.RandToken(),
	})
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	tx.Commit()
	json.NewEncoder(g.Writer).Encode(domain.UserOutputMapping(dbUser))
}
