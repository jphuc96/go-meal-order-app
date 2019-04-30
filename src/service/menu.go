package service

import (
	"database/sql"
	"time"

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

func (s *Service) GetLatestMenu(tx *sql.Tx) (*models.Menu, error) {
	m, _ := s.Store.MenuStore.GetLatestMenu(tx)
	createAt := m.CreatedAt.Truncate(24 * time.Hour)
	now := time.Now().Truncate(24 * time.Hour)

	if !now.Equal(createAt) {
		return nil, nil
	}
	return m, nil
}
