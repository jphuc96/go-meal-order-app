package service

import (
	"errors"
	"regexp"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

var (
	re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func (s *Service) CreateUser(p *domain.UserInput) (*models.User, error) {

	if !(re.MatchString(p.Email)) {
		return nil, errors.New("Invalid email address")
	}

	u, err := s.Store.UserStore.Find(p)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return s.Store.UserStore.Create(p)
	}
	return nil, errors.New("User existed")
}

func (s *Service) GetAllUser() ([]*models.User, error) {
	u, err := s.Store.UserStore.FindAll()
	return u, err
}
