package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetOrdersOfUser(g *gin.Context) {
	menuID := g.Param("MenuID")
	userID := g.Param("UserID")
	items, err := c.service.GetOrdersByMenuAndUser(menuID, userID)
	if err != nil {
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	itemsRes := make([]domain.OrderItem, len(items))
	for i, item := range items {
		itemsRes[i].ID = item.ID
		itemsRes[i].Name = item.ItemName
	}

	json.NewEncoder(g.Writer).Encode(&domain.OrderResp{
		Items: itemsRes,
	})
}

func (c *CoreHandler) CreateOrModifyOrder(g *gin.Context) {
	if domain.MenuStatus == domain.MenuClose {
		c.HandleHTTPError(domain.MenuClosed, http.StatusForbidden, g.Writer)
		return
	}

	menuID := g.Param("MenuID")
	userID := g.Param("UserID")

	var newItems domain.OrderReq
	d := json.NewDecoder(g.Request.Body)

	err := d.Decode(&newItems)
	if err != nil {
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}
	defer g.Request.Body.Close()

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	// Check if all requested items exist
	invalidItemID := make([]int, 0)
	for _, item := range newItems.ItemIDs {
		exist, err := c.service.CheckItemExist(tx, item)
		if err != nil {
			c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
			return
		}

		if !exist {
			invalidItemID = append(invalidItemID, item)
		}
	}

	if len(invalidItemID) != 0 {
		c.HandleHTTPError(fmt.Errorf("invalid value: %v", invalidItemID), http.StatusBadRequest, g.Writer)
		return
	}

	_, err = c.deleteAllOrdersByMenuAndUser(tx, menuID, userID)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}
	tx.Commit()

	tx, err = c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	// var orderResp domain.OrderResp
	for _, itemID := range newItems.ItemIDs {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
			return
		}

		_, err = c.service.AddOrder(tx, &domain.OrderInput{
			UserID: userIDInt,
			ItemID: itemID,
		})
		if err != nil {
			tx.Rollback()
			c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
			return
		}
	}

	tx.Commit()
	c.GetOrdersOfUser(g)
}

func (c *CoreHandler) CancelAllOrderOfUser(g *gin.Context) {
	if domain.MenuStatus == domain.MenuClose {
		c.HandleHTTPError(domain.MenuClosed, http.StatusForbidden, g.Writer)
		return
	}

	menuID := g.Param("MenuID")
	userID := g.Param("UserID")

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		c.HandleHTTPError(err, http.StatusInternalServerError, g.Writer)
	}

	delItems, err := c.deleteAllOrdersByMenuAndUser(tx, menuID, userID)
	if err != nil {
		tx.Rollback()
		c.HandleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	var orderResp domain.OrderResp
	for _, item := range delItems {
		orderResp.Items = append(orderResp.Items, domain.OrderItem{
			ID:   item.ID,
			Name: item.ItemName,
		})
	}

	tx.Commit()
	json.NewEncoder(g.Writer).Encode(orderResp)
}

func (c *CoreHandler) deleteAllOrdersByMenuAndUser(tx *sql.Tx, menuID string, userID string) ([]*domain.Item, error) {
	delItems, err := c.service.GetOrdersByMenuAndUser(menuID, userID)
	if err != nil {
		return nil, err
	}

	for _, item := range delItems {
		userIDStr, err := strconv.Atoi(userID)
		if err != nil {
			return nil, err
		}
		err = c.service.DeleteOrder(tx, &domain.OrderInput{
			UserID: userIDStr,
			ItemID: item.ID,
		})
		if err != nil {
			return nil, err
		}
	}

	return delItems, nil
}
