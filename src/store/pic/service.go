package pic

import (
	"database/sql"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type Service interface {
	Add(tx *sql.Tx, p *domain.PICInput) (*models.PeopleInCharge, error)
	GetByMenuID(tx *sql.Tx, menuID int) ([]*models.PeopleInCharge, error)
	Exist(tx *sql.Tx, p *domain.PICInput) (bool, error)
}
