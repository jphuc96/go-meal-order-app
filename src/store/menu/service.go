package menu

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	CheckMenuExist(menuID int) (bool, error)
	FindByID(mn *domain.MenuInput) (*models.Menu, error)
	Create(p *domain.MenuInput) (*models.Menu, error)
	FindByName(menuName string) (*models.Menu, error)
}
