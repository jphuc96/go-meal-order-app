package menu

import (
	"context"
	"database/sql"

<<<<<<< HEAD
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type menuService struct {
	db *sql.DB
}

// NewService ..
func NewService(db *sql.DB) Service {
	return &menuService{
=======
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

//MenuService struct
type MenuService struct {
	db *sql.DB
}

//NewService function
func NewService(db *sql.DB) Service {
	return &MenuService{
>>>>>>> API to create menu
		db: db,
	}
}

<<<<<<< HEAD
func mapMenuInputToModel(m *domain.MenuInput) *models.Menu {
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

func (s *menuService) FindByID(mn *domain.MenuInput) (*models.Menu, error) {
	m := mapMenuInputToModel(mn)
	return models.FindMenu(context.Background(), s.db, m.ID)
=======
//Create function
func (mn *MenuService) Create(p *domain.MenuInput) (*models.Menu, error) {
	m := &models.Menu{
		OwnerID:         p.OwnerID,
		MenuName:        p.MenuName,
		Deadline:        p.Deadline,
		PaymentReminder: p.PaymentReminder,
		Status:          p.Status,
	}
	return m, m.Insert(context.Background(), mn.db, boil.Infer())
}

//Find function to check if menu is existed
func (mn *MenuService) FindByName(menuName string) (*models.Menu, error) {
	return models.Menus(qm.Where("menu_name = ?", menuName)).One(context.Background(), mn.db)
>>>>>>> API to create menu
}
