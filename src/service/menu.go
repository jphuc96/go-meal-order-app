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
	m, _ := s.Store.MenuStore.GetLatestMenu(tx)
	if m == nil {
		return nil, nil
	}
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return nil, err
	}
	createAt := m.CreatedAt.In(loc).Truncate(24 * time.Hour)
	now := time.Now().In(loc).Truncate(24 * time.Hour)

	if !now.Equal(createAt) {
		return nil, nil
	}
	return m, nil
}

func (s *Service) GetMenuByID(tx *sql.Tx, menuID int) (*models.Menu, error) {
	m, _ := s.Store.MenuStore.FindByID(tx, menuID)
	return m, nil
}

func (s *Service) UpdateMenu(tx *sql.Tx, updateMenu *models.Menu) (*models.Menu, error) {
	_, err := updateMenu.Update(context.Background(), tx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return s.Store.MenuStore.FindByID(tx, updateMenu.ID)
}

func (s *Service) HandleMenuDeadline(tx *sql.Tx, menu *models.Menu) error {
	dl := menu.Deadline.Truncate(time.Minute)
	now := time.Now().Truncate(time.Minute)
	menu.Status = domain.MenuOpen
	if now.After(dl) {
		menu.Status = domain.MenuClose
	}
	domain.MenuStatus = menu.Status

	_, err := s.UpdateMenu(tx, menu)
	return err
}
