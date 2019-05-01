package domain

import (
	"errors"
)

var (
	MenuNotExist     = errors.New("menu does not exist")
	UserNotExist     = errors.New("user does not exist")
	ItemNotExist     = errors.New("item does not exist")
	OrderNotExist    = errors.New("order does not exist")
	TxRollBack       = errors.New("transaction failed, rolling back")
	TxCreateFailed   = errors.New("create transaction failed")
	PICExist         = errors.New("PIC already exists")
	InvalidEmail     = errors.New("invalid email format")
	UserExist        = errors.New("user already exists")
	OrderExist       = errors.New("order already exists")
	MenuExist        = errors.New("menu already exists")
	CreateUserFailed = errors.New("create user operation failed")
	UserNotExistInFT = errors.New("user does not exist in Fortress")
	InvalidMenuID    = errors.New("invalid menu_id")
	InvalidItemID    = errors.New("invalid item_id")
	InvalidUserID    = errors.New("invalid user_id")
)

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
