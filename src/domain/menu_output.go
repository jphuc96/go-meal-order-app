package domain

import (
	"git.d.foundation/datcom/backend/models"
)

// GSOUsers is get summary output user handler struct
type GSOUsers struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

// GSOItems is get summary output items handler struct
type GSOItems struct {
	ID       int        `json:"id,omitempty"`
	ItemName string     `json:"item_name,omitempty"`
	User     []GSOUsers `json:"user,omitempty"`
}

type GetSumaryOutput struct {
	Menu *models.Menu `json:"menu,omitempty"`
	Item []GSOItems   `json:"item,omitempty"`
}
