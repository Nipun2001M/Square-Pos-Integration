package parsers

import (
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	qkit "github.com/clubpay/qlubkit-go"
	"squarepos/dto"
)

func ParseUsingSdk(Incoming qlub.SubmitOrderCommand) dto.OrderRequest {
	var RequestPayload dto.OrderRequest
	RequestPayload.LocationId = Incoming.RestaurantUnique
	RequestPayload.Tableid = Incoming.Order.TableID
	for _, Item := range Incoming.Order.Products {
		RequestPayload.BuyItems = append(RequestPayload.BuyItems, dto.BuyItem{
			Quantity: qkit.Float64ToStr(Item.Qty.Value()),
			Name:     Item.ProductName,
			BasePrice: dto.BasePriceMoney{
				Amount:   qkit.StrToInt(Item.UnitPrice),
				Currency: "USD",
			},
		})
	}
	return RequestPayload

}
