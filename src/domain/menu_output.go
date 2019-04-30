package domain

import (
	"git.d.foundation/datcom/backend/models"
)

type OrderUser struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

// GSOItems is get summary output items handler struct
type MenuItem struct {
	ID       int         `json:"id,omitempty"`
	ItemName string      `json:"item_name,omitempty"`
	Users    []OrderUser `json:"users,omitempty"`
}

// menu's people in charge
type MenuPIC struct {
	USerID   int    `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

type RespMenu struct {
	Menu           *models.Menu `json:"menu,omitempty"`
	Items          []MenuItem   `json:"items,omitempty"`
	PeopleInCharge []MenuPIC    `json:"people_in_charge,omitempty"`
}
