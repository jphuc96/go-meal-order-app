package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) AddItemToMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID, err := strconv.Atoi(vars["MenuID"])
	if err != nil {
		handleHTTPError(domain.InvalidMenuID, http.StatusBadRequest, w)
		return
	}

	newItem := &domain.Item{}
	d := json.NewDecoder(r.Body)
	err = d.Decode(&newItem)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}
	defer r.Body.Close()

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	itemResp, err := c.service.AddItemToMenu(tx, newItem.ItemName, menuID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&domain.ItemResp{
		Item: domain.Item{
			ID:       itemResp.ID,
			ItemName: itemResp.ItemName,
			MenuID:   itemResp.MenuID,
		},
	})
}

func (c *CoreHandler) DeleteItemFromMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["ItemID"])
	if err != nil {
		handleHTTPError(domain.InvalidItemID, http.StatusBadRequest, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	delItem, err := c.service.GetItemByID(tx, itemID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	err = c.service.DeleteItem(tx, delItem.ID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	tx.Commit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&domain.ItemResp{
		Item: domain.Item{
			ID:       delItem.ID,
			ItemName: delItem.ItemName,
			MenuID:   delItem.MenuID,
		},
	})
}
