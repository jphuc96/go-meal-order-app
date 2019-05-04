package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetLatestMenu(g *gin.Context) {

	menuResp := &domain.MenuResp{}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	latestMenu := &models.Menu{}
	latestMenu, err = c.service.GetLatestMenu(tx)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}
	// if no menu today
	if latestMenu == nil {
		return
	}

	menuItems, err := c.service.GetAllItemsByMenuID(tx, latestMenu.ID)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	respItems := make([]domain.MenuItem, 0)

	for _, item := range menuItems {
		users, err := c.service.GetOrderUsersByItem(tx, item.ID)
		if err != nil {
			tx.Rollback()
			c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
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
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	respPIC := make([]domain.MenuPIC, 0)
	for _, pic := range menuPIC {
		user, err := c.service.GetUserbyPIC(tx, pic)
		if err != nil {
			tx.Rollback()
			c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
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
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&menuResp)
}

func (c *CoreHandler) CreateMenu(g *gin.Context) {
	d := json.NewDecoder(g.Request.Body)
	defer g.Request.Body.Close()
	menuReq := &domain.MenuReq{}
	err := d.Decode(&menuReq)
	if err != nil {
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	menu, err := c.service.CreateMenu(tx, &menuReq.Menu)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	menuItems := make([]domain.MenuItem, 0)
	for _, itemName := range menuReq.ItemNames {
		item, err := c.service.AddItemToMenu(tx, itemName, menu.ID)
		if err != nil {
			tx.Rollback()
			c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
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
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(menuResp)
}

func (c *CoreHandler) ModifyMenuTime(g *gin.Context) {
	menuID, err := strconv.Atoi(g.Param("MenuID"))
	if err != nil {
		c.HandleHTTPError(domain.InvalidMenuID, http.StatusBadRequest, g.Writer)
		return
	}

	menuTime := &domain.MenuTime{}
	d := json.NewDecoder(g.Request.Body)
	err = d.Decode(&menuTime)
	if err != nil {
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}
	defer g.Request.Body.Close()

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	m, err := c.service.GetMenuByID(tx, menuID)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	if !menuTime.Deadline.IsZero() {
		m.Deadline = menuTime.Deadline
	}
	if !menuTime.PaymentReminder.IsZero() {
		m.PaymentReminder = menuTime.PaymentReminder
	}

	newMenu, err := c.service.UpdateMenu(tx, m)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	tx.Commit()
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&domain.MenuTime{
		Deadline:        newMenu.Deadline,
		PaymentReminder: newMenu.PaymentReminder,
	})
}
