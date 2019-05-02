package service

import (
	"database/sql"
	"math/rand"
	"time"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (s *Service) AddPIC(tx *sql.Tx, p *domain.PICInput) (*models.PeopleInCharge, error) {
	exist, _ := s.Store.PICStore.Exist(tx, p)
	if exist {
		return nil, domain.PICExist
	}

	return s.Store.PICStore.Add(tx, p)
}

func (s *Service) GetPICByMenuID(tx *sql.Tx, menuID int) ([]*models.PeopleInCharge, error) {
	people, err := s.Store.PICStore.GetByMenuID(tx, menuID)
	if err != nil {
		return nil, err
	}

	return people, nil
}

func (s *Service) GetUserbyPIC(tx *sql.Tx, pic *models.PeopleInCharge) (*models.User, error) {
	return s.Store.UserStore.GetByID(tx, pic.UserID)
}

func (s *Service) GetAllOrderUserOfMenu(tx *sql.Tx, menuID int) ([]*models.User, error) {
	items, err := s.Store.ItemStore.GetAllItemsByMenuID(tx, menuID)
	if err != nil {
		return nil, err
	}

	menuOrders := make([]*models.Order, 0)
	for _, item := range items {
		orders, err := s.Store.OrderStore.GetAllOrdersByItemID(tx, item.ID)
		if err != nil {
			return nil, err
		}
		menuOrders = append(menuOrders, orders...)
	}

	orderUserIDs := make([]int, 0)
	for _, order := range menuOrders {
		orderUserIDs = append(orderUserIDs, order.UserID)
	}

	orderUsers := make([]*models.User, 0)
	for _, userID := range s.unique(orderUserIDs) {
		user, err := s.Store.UserStore.GetByID(tx, userID)
		if err != nil {
			return nil, err
		}
		orderUsers = append(orderUsers, user)
	}

	return orderUsers, nil
}

func (s *Service) GenerateRandomPIC(users []*models.User) ([]domain.PICUser, error) {
	picUsers := make([]domain.PICUser, 0)

	switch len(users) {
	case 0:
	case 1, 2:
		for _, user := range users {
			picUsers = append(picUsers, domain.PICUser{
				ID:   user.ID,
				Name: user.Name,
			})
		}
	default:
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
		for i := 0; i < 2; i++ {
			picUsers = append(picUsers, domain.PICUser{
				ID:   users[i].ID,
				Name: users[i].Name,
			})
		}
	}
	return picUsers, nil
}

func (s *Service) unique(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
