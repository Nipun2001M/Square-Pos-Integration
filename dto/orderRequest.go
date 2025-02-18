package dto


type OrderRequest struct {
	LocationId string	`json:"location_id"`
	BuyItems    []struct{
		Quantity string `json:"quantity"`
		Name string		`json:"name"`
		BasePrice BasePriceMoney `json:"base_price_money"`
	}	`json:"line_items"`
}

type BasePriceMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
} 

type Order struct{
	Order OrderRequest `json:"order"`
}




