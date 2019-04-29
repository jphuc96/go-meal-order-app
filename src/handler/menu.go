package handler

import (
	// "fmt"
	"context"
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetLatestMenu(w http.ResponseWriter, r *http.Request) {

	respMenu := &domain.RespMenu{}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	latestMenu, err := c.service.GetLatestMenu(tx)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}
	// if no menu today
	if latestMenu == nil {
		return
	}

	menuItems, err := c.service.GetAllItemsByMenuID(tx, latestMenu.ID)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	respItems := make([]domain.MenuItems, 0)

	for _, item := range menuItems {
		// users, err := c.service.GetOrdersByItem(tx, item.ID)

		respItems = append(respItems, &domain.MenuItems{
			ID:       item.ID,
			ItemName: item.ItemName,
		})
	}

	respMenu.Items = append(respMenu.Items, respItems)
	respMenu.Menu = latestMenu
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&respMenu)
	// json.NewEncoder(w).Encode()
	// mns, err := c.service.Store.ItemStore.GetAllItemsByMenuID(menuresp.Menu.ID)
	// if err != nil {
	// 	handleHTTPError(err, http.StatusInternalServerError, w)
	// 	return
	// }

	// tx, err := c.db.BeginTx(context.Background(), nil)
	// for _, mn := range mns {
	// 	orders, err := c.service.Store.OrderStore.GetAllOrdersByItemID(tx, mn.ID)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		handleHTTPError(err, http.StatusInternalServerError, w)
	// 		return
	// 	}

	// 	users := []domain.GSOUser{}
	// 	for _, order := range orders {
	// 		user, err := c.service.Store.UserStore.GetByID(tx, order.UserID)
	// 		if err != nil {
	// 			handleHTTPError(err, http.StatusInternalServerError, w)
	// 			return
	// 		}
	// 		users = append(users, domain.GSOUser{ID: user.ID, UserName: user.Name})
	// 	}
	// 	menuresp.Items = append(menuresp.Items, domain.GSOItem{ID: mn.ID, ItemName: mn.ItemName, Users: users})
	// }
	// tx.Commit()

	// resp, err := json.Marshal(menuresp)
	// if err != nil {
	// 	handleHTTPError(err, http.StatusInternalServerError, w)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// w.Write(resp)
}
