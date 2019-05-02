package menu

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	CheckMenuExist(tx *sql.Tx, menuID int) (bool, error)
	Create(tx *sql.Tx, p *domain.MenuInput) (*models.Menu, error)
	IsMenuNameUnique(tx *sql.Tx, menuName string) (bool, error)
	FindByID(tx *sql.Tx, menuID int) (*models.Menu, error)
	GetLatestMenu(tx *sql.Tx) (*models.Menu, error)
	UpdateMenu(tx *sql.Tx, menuID int, updateMenu *models.Menu) error
}
