package domain

type OrderInput struct {
	UserID int
	ItemID int
}

type OrderJSON struct {
	ItemIDs []int `json:"item_ids"`
}
