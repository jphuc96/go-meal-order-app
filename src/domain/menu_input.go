package domain

import "time"

const (
	MenuOpen  = 1
	MenuClose = 0
)

var MenuStatus = MenuOpen

// MenuInput ..
type MenuInput struct {
	OwnerID         int       `json:"owner_id"`
	MenuName        string    `json:"name"`
	Deadline        time.Time `json:"deadline"`
	PaymentReminder time.Time `json:"payment_reminder"`
	Status          int       `json:"status"`
}

type MenuReq struct {
	Menu      MenuInput `json:"menu"`
	ItemNames []string  `json:"item_names"`
}

type MenuTime struct {
	Deadline        time.Time `json:"deadline"`
	PaymentReminder time.Time `json:"payment_reminder"`
}
