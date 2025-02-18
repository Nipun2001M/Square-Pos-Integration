package response

type OrderResponse struct{
	Order order	`json:"order"`
}

type order struct{
	Created string 		`json:"created_at"`
	ID string 			`json:"id"`
	LineItems	[]Item 	`json:"line_items"` 
	LocationId string 	`json:"location_id"` 
	NetAmount Money		`json:"net_amount_due_money"`
	TotalMoney Money 	`json:"total_money"`
	
}
type Item struct{
	BasePrice Money		`json:"base_price_money"`
	Name string			`json:"name"`
	Quantity string 	`json:"quantity"`
	TotalDiscount Money	`json:"total_discount_money"`
	TotalMoney Money	`json:"total_money"`

}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}