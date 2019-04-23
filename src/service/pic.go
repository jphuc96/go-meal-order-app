package service

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

func (s *Service) AddPIC(p *domain.PICInput) (*models.PeopleInCharge, error) {
	exist, _ := s.Store.PICStore.Exist(p)
	if exist {
		return nil, domain.PICExist
	}

	return s.Store.PICStore.Add(p)
}

func (s *Service) GetPICByMenuID(menuID int) ([]*models.PeopleInCharge, error) {
	people, err := s.Store.PICStore.GetByMenuID(menuID)
	if err != nil {
		return nil, err
	}

	return people, nil
}
