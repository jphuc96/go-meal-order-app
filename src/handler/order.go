package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetOrdersOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["MenuID"]
	userID := vars["UserID"]
	items, err := c.service.GetOrdersByMenuAndUser(menuID, userID)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	itemsRes := make([]domain.OrderItem, len(items))
	for i, item := range items {
		itemsRes[i].ID = item.ID
		itemsRes[i].Name = item.ItemName
	}

	json.NewEncoder(w).Encode(&domain.OrderResp{
		Items: itemsRes,
	})
}

func (c *CoreHandler) CreateOrModifyOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["MenuID"]
	userID := vars["UserID"]

	var newItems domain.OrderReq
	d := json.NewDecoder(r.Body)
	err := d.Decode(&newItems)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	// Check if all requested items exist
	invalidItemID := make([]int, 0)
	for _, item := range newItems.ItemIDs {
		exist, err := c.service.CheckItemExist(item)
		if err != nil {
			handleHTTPError(err, http.StatusInternalServerError, w)
			return
		}

		if !exist {
			invalidItemID = append(invalidItemID, item)
		}
	}

	if len(invalidItemID) != 0 {
		handleHTTPError(fmt.Errorf("invalid value: %v", invalidItemID), http.StatusBadRequest, w)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
	}

	_, err = c.deleteAllOrdersByMenuAndUser(tx, menuID, userID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}
	tx.Commit()

	tx, err = c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
	}

	// var orderResp domain.OrderResp
	for _, itemID := range newItems.ItemIDs {
		userIDInt, err := strconv.Atoi(userID)
		if err != nil {
			handleHTTPError(err, http.StatusBadRequest, w)
			return
		}

		_, err = c.service.AddOrder(tx, &domain.OrderInput{
			UserID: userIDInt,
			ItemID: itemID,
		})
		if err != nil {
			tx.Rollback()
			handleHTTPError(err, http.StatusBadRequest, w)
			return
		}
	}

	tx.Commit()
	c.GetOrdersOfUser(w, r)
}

func (c *CoreHandler) CancelAllOrderOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["MenuID"]
	userID := vars["UserID"]

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, w)
	}

	delItems, err := c.deleteAllOrdersByMenuAndUser(tx, menuID, userID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, w)
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
	json.NewEncoder(w).Encode(orderResp)
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
