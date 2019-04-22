package user

import (
	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Create(p *domain.CreateUserInput) (*models.User, error)
	Find(p *domain.CreateUserInput) (*models.User, error)
	FindAll() ([]*models.User, error)
}
