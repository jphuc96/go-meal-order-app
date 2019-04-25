package service

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//CreateMenu function
func (s *Service) CreateMenu(p *domain.MenuInput) (*models.Menu, error) {
	exist, err := s.Store.MenuStore.IsMenuNameUnique(p.MenuName)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, domain.MenuExist
	}
	return s.Store.MenuStore.Create(p)
}