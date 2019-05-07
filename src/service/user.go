package service

import (
	"database/sql"
	"regexp"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

var (
	re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (s *Service) CreateUser(tx *sql.Tx, p *domain.UserInput) (*models.User, error) {

	if !(re.MatchString(p.Email)) {
		return nil, domain.InvalidEmail
	}

	exist, err := s.Store.UserStore.Exist(p)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, domain.UserExist
	}
	return s.Store.UserStore.Create(tx, p)
}

func (s *Service) GetAllUser() ([]*models.User, error) {
	u, err := s.Store.UserStore.FindAll()
	return u, err
}

func (s *Service) GetUserByEmail(tx *sql.Tx, m string) (*models.User, error) {
	return s.Store.UserStore.Find(tx, &domain.UserInput{
		Email: m,
	})
}

func (s *Service) GetUserByToken(tx *sql.Tx, tok string) (*models.User, error) {
	return s.Store.UserStore.GetByToken(tx, tok)
}

func (s *Service) UpdateUserToken(tx *sql.Tx, p *domain.UserInput, newToken string) error {
	return s.Store.UserStore.UpdateToken(tx, p, newToken)
}
