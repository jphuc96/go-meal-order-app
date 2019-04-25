package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetLatestMenu(w http.ResponseWriter, r *http.Request) {

	menuresp := &domain.GetSumaryOutput{}
	var err error
	menuresp.Menu, err = c.service.Store.MenuStore.GetLatestMenu()
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	mns, err := c.service.Store.ItemStore.GetAllItemsByMenuID(menuresp.Menu.ID)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	for _, mn := range mns {
		orders, err := c.service.Store.OrderStore.GetAllOrdersByItemID(tx, mn.ID)
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}

		users := []domain.GSOUsers{}
		for _, order := range orders {
			user, err := c.service.Store.UserStore.GetByID(tx, order.UserID)
			if err != nil {
				handleHTTPError(err, http.StatusInternalServerError, w)
				return
			}
			users = append(users, domain.GSOUsers{ID: user.ID, UserName: user.Name})
		}
		menuresp.Item = append(menuresp.Item, domain.GSOItems{ID: mn.ID, ItemName: mn.ItemName, User: users})
	}
	tx.Commit()

	resp, err := json.Marshal(menuresp)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}
	w.Write(resp)
}
