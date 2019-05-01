package service

import (
	"database/sql"
	"time"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//CreateMenu function
func (s *Service) CreateMenu(tx *sql.Tx, p *domain.MenuInput) (*models.Menu, error) {
	exist, err := s.Store.MenuStore.IsMenuNameUnique(tx, p.MenuName)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, domain.MenuExist
	}
	return s.Store.MenuStore.Create(tx, p)
}

func (s *Service) GetLatestMenu(tx *sql.Tx) (*models.Menu, error) {
	m, err := s.Store.MenuStore.GetLatestMenu(tx)
	if err != nil {
		return nil, err
	}
	createAt := m.CreatedAt.Truncate(24 * time.Hour)
	now := time.Now().Truncate(24 * time.Hour)

	if !now.Equal(createAt) {
		return nil, nil
	}
	return m, nil
}
