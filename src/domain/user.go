package domain

import (
	"git.d.foundation/datcom/backend/models"
)

type UserInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

type UserOutput struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Token       string `json:"access_token,omitempty"`
	DeviceToken string `json:"device_token,omitempty"`
	DeviceType  string `json:"device_type,omitempty"`
}

func UserOutputMapping(u *models.User) UserOutput {
	return UserOutput{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Token: u.Token,
	}
}
