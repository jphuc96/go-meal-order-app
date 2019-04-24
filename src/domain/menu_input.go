package domain

import "time"

// MenuInput ..
type CreateMenuInput struct {
	ID              int
	OwnerID         int
	MenuName        string
	Deadline        time.Time
	PaymentReminder time.Time
	Status          int
	ItemName        []string
}
