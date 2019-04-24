package menu

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

//MenuService struct
type menuService struct {
	db *sql.DB
}

//NewService function
func NewService(db *sql.DB) Service {
	return &menuService{
		db: db,
	}
}

func mapMenuInputToModel(m *domain.CreateMenuInput) *models.Menu {
	return &models.Menu{
		ID:              m.ID,
		OwnerID:         m.OwnerID,
		MenuName:        m.MenuName,
		Deadline:        m.Deadline,
		PaymentReminder: m.PaymentReminder,
		Status:          m.Status,
	}
}

func (s *menuService) CheckMenuExist(menuID int) (bool, error) {
	return models.Menus(
		qm.Where("id = ?", menuID),
	).Exists(context.Background(), s.db)
}

//Create function
func (mn *menuService) Create(p *domain.CreateMenuInput) (*models.Menu, error) {
	m := &models.Menu{
		OwnerID:         p.OwnerID,
		MenuName:        p.MenuName,
		Deadline:        p.Deadline,
		PaymentReminder: p.PaymentReminder,
		Status:          p.Status,
	}
	return m, m.Insert(context.Background(), mn.db, boil.Infer())
}

func (mn *menuService) IsMenuNameUnique(menuName string) (bool, error) {
	return models.Menus(qm.Where("menu_name=?", menuName)).Exists(context.Background(), mn.db)
}
