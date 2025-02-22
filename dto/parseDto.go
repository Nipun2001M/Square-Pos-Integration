package dto

// struct to take response came from squre pos
type OrderResponse struct {
	Order order `json:"order"`
}

type order struct {
	Created    string    `json:"created_at"`
	ID         string    `json:"id"`
	LineItems  []Item    `json:"line_items"`
	TotalMoney Money     `json:"total_money"`
	LocationId string    `json:"location_id"`
	State      string    `json:"state"`
	Tax        Money     `json:"total_tax_money"`
	TrimPrefix Money     `json:"total_tip_money"`
	NetAmout   NetAmount `json:"net_amounts"`
	Due        Money     `json:"net_amount_due_money"`
	Tableid    string    `json:"reference_id"`
}
type Item struct {
	ID            string     `json:"uid"`
	BasePrice     Money      `json:"base_price_money"`
	Name          string     `json:"name"`
	Quantity      string     `json:"quantity"`
	TotalDiscount Money      `json:"total_discount_money"`
	TotalMoney    Money      `json:"total_money"`
	Modifiers     []Modifier `json:"modifiers"`
}

type Modifier struct {
	Name      string `json:"name"`
	BasePrice Money  `json:"base_price_money"`
	Quantity  string `json:"quantity"`
	Amount    Money  `json:"total_price_money"`
}
type NetAmount struct {
	ServiceCharge Money `json:"service_charge_money"`
	Discount      Money `json:"discount_money"`
	Tips          Money `json:"tip_money"`
}

// structs to format response to send
type OrderResponseOut struct {
	Id       string         `json:"id"`
	OpenedAt string         `json:"opened_at"`
	Isclosed bool           `json:"is_closed"`
	Table    string         `json:"table"`
	Items    []ItemFormated `json:"items"`
	Totals   Totals         `json:"totals"`
}

type ItemFormated struct {
	Name      string      `json:"name"`
	Comment   string      `json:"comment"`
	UnitPrice int         `json:"unit_price"`
	Quantity  int         `json:"quantity"`
	Amout     int         `json:"amount"`
	Discounts Discount    `json:"discounts"`
	Modifiers []Modifiers `json:"modifiers"`
}
type Discount struct {
	Name         string `json:"name"`
	IsPercentage bool   `json:"is_percentage"`
	Value        int    `json:"value"`
	Amount       int    `json:"amount"`
}

type Modifiers struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unit_price"`
	Quantity  string `json:"quantity"`
	Amount    int    `json:"amount"`
}
type Totals struct {
	Discounts     int `json:"discounts"`
	Due           int `json:"due"`
	Tax           int `json:"tax"`
	ServiceCharge int `json:"service_charge"`
	Paid          int `json:"paid"`
	Tips          int `json:"tips"`
	Total         int `json:"total"`
}
