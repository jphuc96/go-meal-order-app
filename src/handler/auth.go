package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/src/domain"
)

const (
	oauthGoogleUrlAPI = "https://wwg.Writer.googleapis.com/oauth2/v2/userinfo?access_token="
)

func (c *CoreHandler) GoogleLogin(g *gin.Context) {
	State = c.service.RandToken()
	session := sessions.Default(g)
	session.Set("state", State)
	err := session.Save()
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	url := OAuthConfig.AuthCodeURL(State)

	g.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *CoreHandler) GoogleLogout(g *gin.Context) {
}

func (c *CoreHandler) GoogleOauthCallback(g *gin.Context) {
	session := sessions.Default(g)

	getState := session.Get("state")
	queryState := g.Request.URL.Query().Get("state")
	if getState != queryState {
		handleHTTPError(domain.InvalidOAuthState, http.StatusUnauthorized, g.Writer)
		return
	}

	code := g.Request.URL.Query().Get("code")
	data, err := c.service.GetUserDataFromGoogle(OAuthConfig, code)
	if getState != queryState {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	googleUser := domain.GoogleUser{}
	if err = json.Unmarshal(data, &googleUser); err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	session.Set("user-email", googleUser.Email)
	err = session.Save()
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	json.NewEncoder(g.Writer).Encode(googleUser)
	// tx, err := c.db.BeginTx(context.Background(), nil)
	// if err != nil {
	// 	handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	// 	return
	// }
	// // Check user in db
	// dbUser, _ := c.service.GetUserByEmail(tx, googleUser.Email)
	// if dbUser != nil {
	// 	if newToken[0] != dbUser.Token {
	// 		// If logged out or login in new device, update new token to record
	// 		err := c.service.UpdateUserToken(tx, &domain.CreateUserInput{
	// 			Email: dbUser.Email,
	// 			Token: dbUser.Token,
	// 		}, newToken[0])

	// 		if err != nil {
	// 			tx.Rollback()
	// 			handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	// 			return
	// 		}
	// 		dbUser, err = c.service.GetUserByEmail(tx, googleUser.Email)
	// 		if err != nil {
	// 			tx.Rollback()
	// 			handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	// 			return
	// 		}

	// 	}
	// 	tx.Commit()
	// 	json.NewEncoder(g.Writer).Encode(domain.UserOutputMapping(dbUser))
	// 	return
	// }

	// // if user is new, check with Fortress before decide to create user or not
	// ftUser, err := c.service.FortressVerify(googleUser.Email)
	// if err != nil {
	// 	handleHTTPError(err, http.StatusBadRequest, g.Writer)
	// 	return
	// }

	// tx, err = c.db.BeginTx(context.Background(), nil)
	// if err != nil {
	// 	handleHTTPError(err, http.StatusInternalServerError, g.Writer)
	// 	return
	// }

	// dbUser, err = c.service.CreateUser(tx, &domain.CreateUserInput{
	// 	Name:  ftUser.Name,
	// 	Email: ftUser.Email,
	// 	Token: newToken[0],
	// })
	// if err != nil {
	// 	tx.Rollback()
	// 	handleHTTPError(err, http.StatusBadRequest, g.Writer)
	// 	return
	// }

	// tx.Commit()
	// json.NewEncoder(g.Writer).Encode(domain.UserOutputMapping(dbUser))
}
