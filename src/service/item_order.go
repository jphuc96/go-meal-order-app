package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// DeleteItem ..
func (s *Service) DeleteItem(i *domain.Item) (*models.Item, error) {

	exists, err := s.Store.ItemStore.CheckItemExist(i.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("Item does not exist")
	}

	it, err := s.Store.ItemStore.FindByID(i.ID)
	if err != nil {
		return nil, err
	}

	exists, err = s.Store.OrderStore.CheckOrderExistByItemID(i.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		orders, err := s.Store.OrderStore.GetAllOrdersByItemID(i.ID)
		if err != nil {
			return nil, err
		}
		for _, o := range orders {
			err := s.Store.OrderStore.DeleteOrder(o)
			if err != nil {
				return nil, err
			}
		}
	}

	err = s.Store.ItemStore.Delete(it)
	if err != nil {
		return nil, err
	}

	return it, err
}

func (s *Service) CheckItemExist(itemID int) (bool, error) {
	return s.Store.ItemStore.CheckItemExist(itemID)
}

func (s *Service) GetItemByID(itemID int) (*models.Item, error) {
	return s.Store.ItemStore.FindByID(itemID)
}
