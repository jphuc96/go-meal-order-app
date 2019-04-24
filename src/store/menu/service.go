package menu

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

// Service ..
type Service interface {
	CheckMenuExist(menuID int) (bool, error)
	Create(p *domain.CreateMenuInput) (*models.Menu, error)
	IsMenuNameUnique(menuName string) (bool, error)
}
