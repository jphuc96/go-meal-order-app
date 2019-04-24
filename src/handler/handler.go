package handler

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/service"
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
