package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
)

// DeleteItem ..
func (s *Service) DeleteItem(tx *sql.Tx, itemID int) error {
	item, err := s.Store.ItemStore.FindByID(tx, itemID)
	if err != nil {
		return err
	}

	orders, err := s.Store.OrderStore.GetAllOrdersByItemID(tx, itemID)
	if err != nil {
		return err
	}
	for _, o := range orders {
		err := s.Store.OrderStore.DeleteOrder(tx, o)
		if err != nil {
			return err
		}
	}

	return s.Store.ItemStore.Delete(tx, item)
}

func (s *Service) CheckItemExist(tx *sql.Tx, itemID int) (bool, error) {
	return s.Store.ItemStore.CheckItemExist(tx, itemID)
}

func (s *Service) GetItemByID(tx *sql.Tx, itemID int) (*models.Item, error) {
	return s.Store.ItemStore.FindByID(tx, itemID)
}
