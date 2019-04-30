package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//CreateMenu function
func (s *Service) CreateMenu(p *domain.CreateMenuInput) (*models.Menu, error) {
	exist, err := s.Store.MenuStore.IsMenuNameUnique(p.MenuName)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New(domain.MenuExist)
	}
	return s.Store.MenuStore.Create(p)
}
