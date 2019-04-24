package order

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error)
	Delete(tx *sql.Tx, o *domain.OrderInput) error
	Exist(o *domain.OrderInput) (bool, error)
	Get(menuID string, userID string) ([]*domain.Item, error)
	DeleteOrder(o *models.Order) error
	CheckOrderExistByItemID(ItemID int) (bool, error)
	GetAllOrdersByItemID(ItemID int) ([]*models.Order, error)
	GetOrderByOrderInput(o *domain.OrderInput) (*models.Order, error)
}
