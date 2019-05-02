package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetPeopleInCharge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID, err := strconv.Atoi(vars["MenuID"])
	if err != nil {
		handleHTTPError(domain.InvalidMenuID, http.StatusBadRequest, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	users, err := c.service.GetAllOrderUserOfMenu(tx, menuID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	picUsers, err := c.service.GenerateRandomPIC(users)
	for _, picUser := range picUsers {
		_, err := c.service.AddPIC(tx, &domain.PICInput{
			MenuID: menuID,
			UserID: picUser.ID,
		})
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&domain.PICResp{
		Users: picUsers,
	})
}
