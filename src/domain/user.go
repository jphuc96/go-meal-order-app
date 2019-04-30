package domain

import (
	"git.d.foundation/datcom/backend/models"
)

type CreateUserInput struct {
	Name     string `json:"name,omitempty"`
	GoogleID string `json:"google_id,omitempty"`
	Email    string `json:"email,omitempty"`
	Token    string `json:"token,omitempty"`
}

type CreateaUserOutput struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Token       string `json:"token,omitempty"`
	DeviceToken string `json:"device_token,omitempty"`
	DeviceType  string `json:"device_type,omitempty"`
}

func UserOutputMapping(u *models.User) CreateaUserOutput {
	return CreateaUserOutput{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Token: u.Token,
	}
}
