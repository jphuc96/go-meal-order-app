package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) AddItemToMenu(g *gin.Context) {
	menuID, err := strconv.Atoi(g.Param("MenuID"))
	if err != nil {
		handleHTTPError(domain.InvalidMenuID, http.StatusBadRequest, g.Writer)
		return
	}

	newItem := &domain.Item{}
	d := json.NewDecoder(g.Request.Body)
	err = d.Decode(&newItem)
	if err != nil {
		handleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}
	defer g.Request.Body.Close()

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	itemResp, err := c.service.AddItemToMenu(tx, newItem.ItemName, menuID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	tx.Commit()
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&domain.ItemResp{
		Item: domain.Item{
			ID:       itemResp.ID,
			ItemName: itemResp.ItemName,
			MenuID:   itemResp.MenuID,
		},
	})
}

func (c *CoreHandler) DeleteItemFromMenu(g *gin.Context) {
	itemID, err := strconv.Atoi(g.Param("ItemID"))
	if err != nil {
		handleHTTPError(domain.InvalidItemID, http.StatusBadRequest, g.Writer)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	delItem, err := c.service.GetItemByID(tx, itemID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusBadRequest, g.Writer)
		return
	}

	err = c.service.DeleteItem(tx, delItem.ID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	tx.Commit()
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&domain.ItemResp{
		Item: domain.Item{
			ID:       delItem.ID,
			ItemName: delItem.ItemName,
			MenuID:   delItem.MenuID,
		},
	})
}
