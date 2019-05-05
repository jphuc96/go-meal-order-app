package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type userService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &userService{
		db: db,
	}
}

func userInputToDBMapping(p *domain.UserInput) *models.User {
	return &models.User{
		Name:  p.Name,
		Email: p.Email,
		Token: p.Token,
	}
}

func (us *userService) Create(tx *sql.Tx, p *domain.UserInput) (*models.User, error) {
	u := userInputToDBMapping(p)
	err := u.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		return nil, errors.New("Insert failed")
	}
	return us.Find(tx, p)
}

func (us *userService) Exist(p *domain.UserInput) (bool, error) {
	return models.Users(qm.Where("email=?", p.Email)).Exists(context.Background(), us.db)
}

func (us *userService) Find(tx *sql.Tx, p *domain.UserInput) (*models.User, error) {
	user, err := models.Users(qm.Where("email=?", p.Email)).One(context.Background(), tx)
	return user, err
}

func (us *userService) FindAll() ([]*models.User, error) {
	users, err := models.Users().All(context.Background(), us.db)
	return users, err
}

func (us *userService) UpdateToken(tx *sql.Tx, p *domain.UserInput, newToken string) error {
	user, err := models.Users(qm.Where("email=? AND token=?", p.Email, p.Token)).One(context.Background(), us.db)
	if err != nil {
		return err
	}
	user.Token = newToken
	_, err = user.Update(context.Background(), tx, boil.Infer())
	return err
}

func (us *userService) GetByID(tx *sql.Tx, userID int) (*models.User, error) {
	return models.FindUser(context.Background(), tx, userID)
}

func (us *userService) ExistByEmailAndToken(email, token string) (bool, error) {
	return models.Users(qm.Where("email=? AND token=?", email, token)).Exists(context.Background(), us.db)
}
