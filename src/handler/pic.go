package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"git.d.foundation/datcom/backend/src/domain"
)

func (c *CoreHandler) GetPeopleInCharge(g *gin.Context) {
	menuID, err := strconv.Atoi(g.Param("MenuID"))
	if err != nil {
		handleHTTPError(domain.InvalidMenuID, http.StatusBadRequest, g.Writer)
		return
	}

	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
		return
	}

	users, err := c.service.GetAllOrderUserOfMenu(tx, menuID)
	if err != nil {
		tx.Rollback()
		handleHTTPError(err, http.StatusInternalServerError, g.Writer)
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
			handleHTTPError(err, http.StatusInternalServerError, g.Writer)
			return
		}
	}

	tx.Commit()
	g.Writer.WriteHeader(http.StatusOK)
	json.NewEncoder(g.Writer).Encode(&domain.PICResp{
		Users: picUsers,
	})
}
