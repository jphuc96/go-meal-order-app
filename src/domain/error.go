package domain

var (
	MenuNotExist       = "menu does not exist"
	UserNotExist       = "user does not exist"
	ItemNotExist       = "item does not exist"
	OrderNotExist      = "order does not exist"
	TxRollBack         = "transaction failed, rolling back"
	TxCreateFailed     = "create transaction failed"
	PICExist           = "PIC is already exist"
	InvalidEmailFormat = "invalid email format"
	UserExist          = "user is already exist"
	OrderExist         = "order is already exist"
	MenuExist          = "menu is already exist"
	CreateUserFailed   = "create user operation is failed"
)

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
