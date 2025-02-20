package dto

type OrderResponse struct {
	Order order `json:"order"`
}

type order struct {
	Created    string `json:"created_at"`
	ID         string `json:"id"`
	LineItems  []Item `json:"line_items"`
	LocationId string `json:"location_id"`
	TotalMoney Money  `json:"total_money"`
	State      string `json:"state"`
	Tax		   	Money	`json:"total_tax_money"`
	TrimPrefix	Money	`json:"total_tip_money"`
}
type Item struct {
	BasePrice     Money  `json:"base_price_money"`
	Name          string `json:"name"`
	Quantity      string `json:"quantity"`
	TotalDiscount Money  `json:"total_discount_money"`
	TotalMoney    Money  `json:"total_money"`
}
