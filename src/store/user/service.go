package user

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Create(tx *sql.Tx, p *domain.UserInput) (*models.User, error)
	Find(tx *sql.Tx, p *domain.UserInput) (*models.User, error)
	Exist(p *domain.UserInput) (bool, error)
	FindAll() ([]*models.User, error)
	UpdateToken(tx *sql.Tx, p *domain.UserInput, newToken string) error
	GetByID(tx *sql.Tx, userID int) (*models.User, error)
	ExistByToken(token string) (bool, error)
	GetByToken(tx *sql.Tx, tok string) (*models.User, error)
}
