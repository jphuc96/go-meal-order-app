package domain

type PICInput struct {
	UserID int `json:"user_id,omitempty"`
	MenuID int `json:"menu_id,omitempty"`
}

type PICResp struct {
	Users []PICUser `json:"users"`
}

type PICUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
