package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// DeleteItem ..
func (s *Service) DeleteItem(tx *sql.Tx, i *domain.Item) (*models.Item, error) {

	exists, err := s.Store.ItemStore.CheckItemExist(tx, i.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, domain.ItemNotExist
	}

	it, err := s.Store.ItemStore.FindByID(tx, i.ID)
	if err != nil {
		return nil, err
	}

	exists, err = s.Store.OrderStore.CheckOrderExistByItemID(tx, i.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		orders, err := s.Store.OrderStore.GetAllOrdersByItemID(tx, i.ID)
		if err != nil {
			return nil, err
		}
		for _, o := range orders {
			err := s.Store.OrderStore.DeleteOrder(tx, o)
			if err != nil {
				return nil, err
			}
		}
	}

	err = s.Store.ItemStore.Delete(tx, it)
	if err != nil {
		return nil, err
	}

	return it, err
}

func (s *Service) CheckItemExist(tx *sql.Tx, itemID int) (bool, error) {
	return s.Store.ItemStore.CheckItemExist(tx, itemID)
}

func (s *Service) GetItemByID(tx *sql.Tx, itemID int) (*models.Item, error) {
	return s.Store.ItemStore.FindByID(tx, itemID)
}
