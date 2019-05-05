package handler

import (
	"database/sql"
	"os"

	"golang.org/x/oauth2"

	"git.d.foundation/datcom/backend/src/service"
)

var (
	State       string
	OAuthConfig *oauth2.Config
)

type CoreHandler struct {
	service *service.Service
	db      *sql.DB
}

func NewCoreHandler(service *service.Service, db *sql.DB) *CoreHandler {
	OAuthConfig = service.ConfigureOAuth2(
		os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		os.Getenv("GOOGLE_REDIRECT_URL"),
	)

	return &CoreHandler{
		service: service,
		db:      db,
	}
}
