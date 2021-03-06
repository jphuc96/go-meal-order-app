package item

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type itemService struct {
	db *sql.DB
}

// NewService ..
func NewService(db *sql.DB) Service {
	return &itemService{
		db: db,
	}
}
func mapItemInputToModel(i *domain.Item) *models.Item {
	return &models.Item{
		ID:       i.ID,
		ItemName: i.ItemName,
		MenuID:   i.MenuID,
	}
}

func (s *itemService) Add(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
	item := mapItemInputToModel(i)
	return item, item.Insert(context.Background(), tx, boil.Infer())
}

func (s *itemService) FindByID(tx *sql.Tx, itemID int) (*models.Item, error) {
	return models.Items(qm.Where("id=?", itemID)).One(context.Background(), tx)
}

func (s *itemService) Delete(tx *sql.Tx, i *models.Item) error {
	_, err := i.Delete(context.Background(), tx)
	return err
}

func (s *itemService) CheckItemExist(tx *sql.Tx, itemID int) (bool, error) {
	b, err := models.Items(qm.Where("id=?", itemID)).Exists(context.Background(), tx)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (s *itemService) GetAllItemsByMenuID(tx *sql.Tx, menuID int) ([]*models.Item, error) {
	return models.Items(qm.Where("menu_id=?", menuID)).All(context.Background(), tx)
}
