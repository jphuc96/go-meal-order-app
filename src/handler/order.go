package handler

import (
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

	itemsRes := make([]int, len(items))
	for i, item := range items {
		itemsRes[i] = item.ID
	}

	json.NewEncoder(w).Encode(&domain.OrderJSON{
		ItemIDs: itemsRes,
	})
}

func (c *CoreHandler) CreateOrModifyOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["MenuID"]
	userID := vars["UserID"]

	oldItems, err := c.service.GetOrdersByMenuAndUser(menuID, userID)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	var newItems domain.OrderJSON
	d := json.NewDecoder(r.Body)
	err = d.Decode(&newItems)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	// for _, item := range newItems.ItemIDs {

	// }

	for _, item := range oldItems {
		userIDStr, _ := strconv.Atoi(userID)
		_ = c.service.DeleteOrder(&domain.OrderInput{
			UserID: userIDStr,
			ItemID: item.ID,
		})
	}

	var orderRes domain.OrderJSON
	for _, item := range newItems.ItemIDs {
		userIDStr, _ := strconv.Atoi(userID)
		orders, err := c.service.AddOrder(&domain.OrderInput{
			UserID: userIDStr,
			ItemID: item,
		})
		if err != nil {
			fmt.Println(err)
		} else {
			orderRes.ItemIDs = append(orderRes.ItemIDs, orders.ItemID)
		}
	}

	json.NewEncoder(w).Encode(orderRes)

}

func (c *CoreHandler) CancelAllOrderOfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menuID := vars["MenuID"]
	userID := vars["UserID"]

	delItems, err := c.service.GetOrdersByMenuAndUser(menuID, userID)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, w)
		return
	}

	for _, item := range delItems {
		userIDStr, _ := strconv.Atoi(userID)
		c.service.DeleteOrder(&domain.OrderInput{
			UserID: userIDStr,
			ItemID: item.ID,
		})
	}

	itemsRes := make([]int, len(delItems))
	for i, item := range delItems {
		itemsRes[i] = item.ID
	}

	json.NewEncoder(w).Encode(&domain.OrderJSON{
		ItemIDs: itemsRes,
	})

}
