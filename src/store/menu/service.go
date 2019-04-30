package menu

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	CheckMenuExist(menuID int) (bool, error)
	Create(p *domain.MenuInput) (*models.Menu, error)
	IsMenuNameUnique(menuName string) (bool, error)
	FindByID(mn *domain.MenuInput) (*models.Menu, error)
	GetLatestMenu(tx *sql.Tx) (*models.Menu, error)
}
