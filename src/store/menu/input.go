package menu

import "time"

type MenuInput struct {
	OwnerID         int
	MenuName        string
	Deadline        time.Time
	PaymentReminder time.Time
	Status          int
	ItemName        []string
}
