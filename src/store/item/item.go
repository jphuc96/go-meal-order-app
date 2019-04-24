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

func (s *itemService) Add(i *domain.Item) (*models.Item, error) {
	item := mapItemInputToModel(i)
	return item, item.Insert(context.Background(), s.db, boil.Infer())
}

func (s *itemService) FindByID(itemID int) (*models.Item, error) {
	return models.FindItem(context.Background(), s.db, itemID)
}

func (s *itemService) Delete(i *models.Item) error {
	_, err := i.Delete(context.Background(), s.db)
	return err
}

func (s *itemService) CheckItemExist(itemID int) (bool, error) {
	b, err := models.Items(qm.Where("id=?", itemID)).Exists(context.Background(), s.db)
	if err != nil {
		return false, err
	}
	return b, nil
}
