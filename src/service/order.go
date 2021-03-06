package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//AddOrder adds new order to resource if order does not exist, otherwise return error
func (s *Service) AddOrder(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
	exist, _ := s.Store.OrderStore.Exist(tx, o)
	if exist {
		return s.Store.OrderStore.GetOrderByOrderInput(tx, o)
	}

	return s.Store.OrderStore.Add(tx, o)
}

//DeleteOrder deletes order from resource if order exists, otherwise return error
func (s *Service) DeleteOrder(tx *sql.Tx, o *domain.OrderInput) error {
	exist, _ := s.Store.OrderStore.Exist(tx, o)
	if !exist {
		return nil
	}

	err := s.Store.OrderStore.Delete(tx, o)

	return err
}

//GetOrdersByUserID gets items in user order by user ID
func (s *Service) GetOrdersByMenuAndUser(menuID string, userID string) ([]*domain.Item, error) {
	items, _ := s.Store.OrderStore.Get(menuID, userID)
	return items, nil
}

func (s *Service) GetOrderUsersByItem(tx *sql.Tx, itemID int) ([]*models.User, error) {
	orderUser := make([]*models.User, 0)
	orders, _ := s.Store.OrderStore.GetAllOrdersByItemID(tx, itemID)

	for _, order := range orders {
		user, err := s.Store.UserStore.GetByID(tx, order.UserID)
		if err != nil {
			return nil, err
		}
		orderUser = append(orderUser, user)
	}

	return orderUser, nil
}
