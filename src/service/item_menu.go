package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// AddItems ..
func (s *Service) AddItems(tx *sql.Tx, items *domain.ItemInput, menuID int) ([]*models.Item, error) {
	var list []*models.Item

	exists, err := s.Store.MenuStore.CheckMenuExist(tx, menuID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.MenuNotExist
	}

	for _, it := range items.Items {
		i := &domain.Item{ItemName: it.ItemName, MenuID: menuID}
		exists, err := s.Store.ItemStore.CheckItemExist(tx, i.ID)
		if err != nil {
			return nil, err
		}
		if exists {
			continue
		}
		resItem, err := s.Store.ItemStore.Add(tx, i)
		if err != nil {
			return nil, err
		}
		list = append(list, resItem)
	}

	return list, err
}

func (s *Service) GetAllItemsByMenuID(tx *sql.Tx, menuID int) ([]*models.Item, error) {
	i, _ := s.Store.ItemStore.GetAllItemsByMenuID(tx, menuID)
	return i, nil
}

func (s *Service) AddItemToMenu(tx *sql.Tx, itemName string, menuID int) (*models.Item, error) {
	exists, err := s.Store.MenuStore.CheckMenuExist(tx, menuID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.MenuNotExist
	}

	return s.Store.ItemStore.Add(tx, &domain.Item{
		ItemName: itemName,
		MenuID:   menuID,
	})
}
