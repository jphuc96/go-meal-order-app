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
	PICExist         = errors.New("PIC is already exist")
	InvalidEmail     = errors.New("invalid email format")
	UserExist        = errors.New("user is already exist")
	OrderExist       = errors.New("order is already exist")
	MenuExist        = errors.New("menu is already exist")
	CreateUserFailed = errors.New("create user operation is failed")
	UserNotExistInFT = errors.New("user does not exist in Fortress")
)

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
