package dto

type OrderRequest struct {
	LocationId string    `json:"location_id"`
	Tableid    string    `json:"reference_id"`
	BuyItems   []BuyItem `json:"line_items"`
}

type BuyItem struct {
	Quantity  string         `json:"quantity"`
	Name      string         `json:"name"`
	BasePrice BasePriceMoney `json:"base_price_money"`
}

type BasePriceMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type Order struct {
	Order OrderRequest `json:"order"`
}
