package order

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error)
	Delete(tx *sql.Tx, o *domain.OrderInput) error
	Exist(tx *sql.Tx, o *domain.OrderInput) (bool, error)
	Get(menuID string, userID string) ([]*domain.Item, error)
	DeleteOrder(tx *sql.Tx, o *models.Order) error
	CheckOrderExistByItemID(tx *sql.Tx, ItemID int) (bool, error)
	GetAllOrdersByItemID(tx *sql.Tx, ItemID int) ([]*models.Order, error)
	GetOrderByOrderInput(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error)
}
