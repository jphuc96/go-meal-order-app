package service

import (
	"errors"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//CreateMenu function
func (s *Service) CreateMenu(p *domain.MenuInput) (*models.Menu, error) {
	menu, err := s.Menu.FindByName(p.MenuName)
	if err != nil {
		return nil, err
	}
	if menu != nil {
		return nil, errors.New("Menu exists")
	}
	return s.Menu.Create(p)
}
