package item

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	Add(tx *sql.Tx, i *domain.Item) (*models.Item, error)
	FindByID(tx *sql.Tx, itemID int) (*models.Item, error)
	Delete(tx *sql.Tx, i *models.Item) error
	CheckItemExist(tx *sql.Tx, itemID int) (bool, error)
}
