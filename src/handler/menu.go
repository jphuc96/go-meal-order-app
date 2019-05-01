package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetLatestMenu(w http.ResponseWriter, r *http.Request) {

	menuResp := &domain.MenuResp{}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	latestMenu := &models.Menu{}
	latestMenu, err = c.service.GetLatestMenu(tx)
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
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	respItems := make([]domain.MenuItem, 0)

	for _, item := range menuItems {
		users, err := c.service.GetOrderUsersByItem(tx, item.ID)
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}

		respUser := make([]domain.OrderUser, 0)
		for _, user := range users {
			respUser = append(respUser, domain.OrderUser{
				ID:       user.ID,
				UserName: user.Name,
			})
		}

		respItems = append(respItems, domain.MenuItem{
			ID:       item.ID,
			ItemName: item.ItemName,
			Users:    respUser,
		})
	}

	menuPIC, err := c.service.GetPICByMenuID(tx, latestMenu.ID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, w)
		return
	}

	respPIC := make([]domain.MenuPIC, 0)
	for _, pic := range menuPIC {
		user, err := c.service.GetUserbyPIC(tx, pic)
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}
		respPIC = append(respPIC, domain.MenuPIC{
			USerID:   user.ID,
			UserName: user.Name,
		})
	}

	menuResp.Items = respItems
	menuResp.Menu = latestMenu
	menuResp.PeopleInCharge = respPIC

	tx.Commit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&menuResp)
}

func (c *CoreHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	menuReq := &domain.MenuReq{}
	err := d.Decode(&menuReq)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
	}

	menu, err := c.service.CreateMenu(tx, &menuReq.Menu)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	menuItems := make([]domain.MenuItem, 0)
	for _, itemName := range menuReq.ItemNames {
		item, err := c.service.AddItemToMenu(tx, itemName, menu.ID)
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}
		menuItems = append(menuItems, domain.MenuItem{
			ID:       item.ID,
			ItemName: item.ItemName,
		})
	}

	menuResp := &domain.MenuResp{
		Menu:  menu,
		Items: menuItems,
	}

	menuResp.Menu.Status = 1

	tx.Commit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menuResp)
}
