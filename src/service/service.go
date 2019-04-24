package service

import (
	"database/sql"

	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/pic"
	"git.d.foundation/datcom/backend/src/store/user"
)

type Service struct {
	Store Store
}

type Store struct {
	UserStore  user.Service
	ItemStore  item.Service
	MenuStore  menu.Service
	OrderStore order.Service
	PICStore   pic.Service
}

func NewService(db *sql.DB) *Service {
	return &Service{
		Store: Store{
			UserStore:  user.NewService(db),
			ItemStore:  item.NewService(db),
			MenuStore:  menu.NewService(db),
			OrderStore: order.NewService(db),
			PICStore:   pic.NewService(db),
		},
	}
}
