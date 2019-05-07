package handler

import (
	"database/sql"

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
	return &CoreHandler{
		service: service,
		db:      db,
	}
}
