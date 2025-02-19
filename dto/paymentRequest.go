package dto


type PaymentRequest struct{
	IdempotencyKey		string				`json:"idempotency_key"`
	AmountMoney			Money 				`json:"amount_money"`
	SourceId 			string				`json:"source_id"` 	   
	OrderId			string				`json:"order_id"` 
	LocationID 			string				`json:"location_id"`

}

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string `json:"currency"`
}