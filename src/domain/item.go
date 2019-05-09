package domain

// ItemInput ..
type ItemInput struct {
	Items []Item
}

// Item ..
type Item struct {
	ID       int    `json:"id"`
	ItemName string `json:"item_name"`
	MenuID   int    `json:"menu_id"`
}

type ItemResp struct {
	Item Item `json:"item"`
}
