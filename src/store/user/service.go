package user

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Create(tx *sql.Tx, p *domain.CreateUserInput) (*models.User, error)
	Find(p *domain.CreateUserInput) (*models.User, error)
	Exist(p *domain.CreateUserInput) (bool, error)
	FindAll() ([]*models.User, error)
	UpdateToken(tx *sql.Tx, p *domain.CreateUserInput, newToken string) error
	GetByID(tx *sql.Tx, userID int) (*models.User, error)
}
