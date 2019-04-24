package domain

type OrderInput struct {
	UserID int
	ItemID int
}

type OrderReq struct {
	ItemIDs []int `json:"item_ids"`
}
type OrderItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type OrderResp struct {
	Items []OrderItem `json:"items"`
}
