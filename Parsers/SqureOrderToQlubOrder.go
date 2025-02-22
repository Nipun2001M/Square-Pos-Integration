package parsers

import (
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	qkit "github.com/clubpay/qlubkit-go"
	"squarepos/dto"
)

func SqureOrderToQlubOrder(Incoming dto.OrderResponse) (qlub.Order, error) {
	Order := qlub.Order{}
	Order.ID = Incoming.Order.ID
	Order.Status = qlub.OrderStatus(Incoming.Order.State)
	Order.TableID = Incoming.Order.Tableid
	Order.Total = qkit.Float64ToStr(Incoming.Order.TotalMoney.Amount)
	Order.QlubDiscounts = append(Order.QlubDiscounts, qlub.QlubDiscounts{
		Amount: qkit.Float64ToStr(Incoming.Order.NetAmout.Discount.Amount),
	})
	for _, item := range Incoming.Order.LineItems {
		var top []qlub.Topping
		for _, modifier := range item.Modifiers {
			top = append(top, qlub.Topping{
				Title:      modifier.Name,
				Quantity:   modifier.Quantity,
				UnitPrice:  qkit.Float64ToStr(modifier.BasePrice.Amount),
				FinalPrice: qkit.Float64ToStr(modifier.Amount.Amount),
			})
		}
		Order.Items = append(Order.Items, qlub.OrderItem{
			Title:      item.Name,
			Quantity:   item.Quantity,
			FinalPrice: qkit.Float64ToStr(item.TotalMoney.Amount),
			UnitPrice:  qkit.Float64ToStr(item.BasePrice.Amount),
			ID:         item.ID,
			Toppings:   top,
		})

	}

	return Order, nil

}
