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
func (mn *menuService) Create(p *domain.MenuInput) (*models.Menu, error) {

	m := mapMenuInputToModel(p)
	tx, err := mn.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, domain.TxCreateFailed
	}
	err = m.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, domain.TxRollBack
	}
	tx.Commit()
	return m, nil
}

func (mn *menuService) IsMenuNameUnique(menuName string) (bool, error) {
	return models.Menus(qm.Where("menu_name=?", menuName)).Exists(context.Background(), mn.db)
}

func (s *menuService) FindByID(mn *domain.MenuInput) (*models.Menu, error) {
	m := mapMenuInputToModel(mn)
	return models.FindMenu(context.Background(), s.db, m.ID)
}

func (s *menuService) GetLatestMenu(tx *sql.Tx) (*models.Menu, error) {
	return models.Menus(qm.OrderBy("created_at DESC")).One(context.Background(), tx)
}
