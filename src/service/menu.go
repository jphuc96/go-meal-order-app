package service

import (
	"context"
	"database/sql"
	"time"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"github.com/volatiletech/sqlboiler/boil"
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

func (s *Service) GetMenuByID(tx *sql.Tx, menuID int) (*models.Menu, error) {
	return s.Store.MenuStore.FindByID(tx, menuID)
}

func (s *Service) UpdateMenu(tx *sql.Tx, updateMenu *models.Menu) (*models.Menu, error) {
	_, err := updateMenu.Update(context.Background(), tx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return s.Store.MenuStore.FindByID(tx, updateMenu.ID)
}
