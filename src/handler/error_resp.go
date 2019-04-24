package handler

import (
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/src/domain"
)

func handleHTTPError(err error, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&domain.ErrorResponse{
		Error: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    statusCode,
			Message: err.Error(),
		},
	})
}
