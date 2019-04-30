package service

import (
	"database/sql"

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
