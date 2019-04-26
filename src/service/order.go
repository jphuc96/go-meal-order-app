package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//AddOrder adds new order to resource if order does not exist, otherwise return error
func (s *Service) AddOrder(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
	exist, _ := s.Store.Order.Exist(o)
	if exist {
		return s.Store.Order.GetOrderByOrderInput(o)
	}

	return s.Store.Order.Add(tx, o)
}

//DeleteOrder deletes order from resource if order exists, otherwise return error
func (s *Service) DeleteOrder(tx *sql.Tx, o *domain.OrderInput) error {
	exist, _ := s.Store.Order.Exist(o)
	if !exist {
		return nil
	}

	err := s.Store.Order.Delete(tx, o)

	return err
}

//GetOrdersByUserID gets items in user order by user ID
func (s *Service) GetOrdersByMenuAndUser(menuID string, userID string) ([]*domain.Item, error) {
	items, err := s.Store.Order.Get(menuID, userID)
	if err != nil {
		return nil, err
	}

	return items, nil
}
