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

func (c *CoreHandler) GoogleLogout(g *gin.Context) {
	tok := g.Request.Header.Get("access_token")

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	u, err := c.service.GetUserByToken(tx, tok)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	err = c.service.UpdateUserToken(tx, &domain.UserInput{
		Email: u.Email,
		Token: u.Token,
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
	idToken := g.Request.Header.Get("id_token")
	if idToken == "" {
		c.HandleHTTPError(domain.NotProvideToken, http.StatusBadRequest, g.Writer)
		return
	}

	// Call google api to verify user
	googleUser, err := c.service.GetGoogleUserInfo(idToken)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	// Check user in db
	dbUser, _ := c.service.GetUserByEmail(tx, googleUser.Email)
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
			dbUser, _ = c.service.GetUserByEmail(tx, googleUser.Email)
		}

		tx.Commit()
		json.NewEncoder(g.Writer).Encode(domain.UserOutputMapping(dbUser))
		return
	}
	// if user is new, check with Fortress before decide to create user or not
	ftUser := &domain.FTResp{}
	_, err = c.service.FortressVerify(googleUser.Email)
	if err != nil {
		c.HandleHTTPError(err, http.StatusUnauthorized, g.Writer)
		return
	}
	if ftUser == nil {
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
