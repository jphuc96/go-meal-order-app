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

func mapMenuInputToModel(m *domain.MenuInput) *models.Menu {
	return &models.Menu{
		OwnerID:         m.OwnerID,
		MenuName:        m.MenuName,
		Deadline:        m.Deadline,
		PaymentReminder: m.PaymentReminder,
		Status:          m.Status,
	}
}

func (ms *menuService) CheckMenuExist(tx *sql.Tx, menuID int) (bool, error) {
	return models.Menus(
		qm.Where("id = ?", menuID),
	).Exists(context.Background(), tx)
}

//Create function
func (ms *menuService) Create(tx *sql.Tx, p *domain.MenuInput) (*models.Menu, error) {
	m := mapMenuInputToModel(p)

	err := m.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (ms *menuService) IsMenuNameUnique(tx *sql.Tx, menuName string) (bool, error) {
	return models.Menus(qm.Where("menu_name=?", menuName)).Exists(context.Background(), tx)
}

func (ms *menuService) FindByID(tx *sql.Tx, menuID int) (*models.Menu, error) {
	return models.FindMenu(context.Background(), tx, menuID)
}

func (ms *menuService) GetLatestMenu(tx *sql.Tx) (*models.Menu, error) {
	return models.Menus(qm.OrderBy("created_at DESC")).One(context.Background(), tx)
}

func (ms *menuService) UpdateMenu(tx *sql.Tx, menuID int, updateMenu *models.Menu) error {
	m, err := ms.FindByID(tx, menuID)
	if err != nil {
		return err
	}

	m = updateMenu
	_, err = m.Update(context.Background(), tx, boil.Infer())
	return err
}
