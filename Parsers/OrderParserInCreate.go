package parsers

import (
	"fmt"
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	qkit "github.com/clubpay/qlubkit-go"
	"squarepos/dto"
	"strconv"
)

func IncomingOrderToSquareParse(Incoming qlub.SubmitOrderCommand) dto.Order {
	var RequestPayload dto.OrderRequest
	RequestPayload.LocationId = Incoming.RestaurantUnique
	RequestPayload.Tableid = Incoming.Order.TableID
	for _, Item := range Incoming.Order.Products {
		var Itemmodifiers []dto.ReqModifier
		for _, modifier := range Item.Modifiers {
			Itemmodifiers = append(Itemmodifiers, dto.ReqModifier{
				Name:     modifier.ProductName,
				Quantity: qkit.Float64ToStr(modifier.Qty.Value()),
				BasePrice: dto.Money{
					Amount:   qkit.StrToFloat64(modifier.UnitPrice),
					Currency: "USD",
				},
			})

		}
		fmt.Println(Item.UnitPrice)
		priceFloat, _ := strconv.ParseFloat(Item.UnitPrice, 64)

		RequestPayload.BuyItems = append(RequestPayload.BuyItems, dto.BuyItem{
			Quantity: qkit.Float64ToStr(Item.Qty.Value()),
			Name:     Item.ProductName,
			BasePrice: dto.BasePriceMoney{
				Amount:   int(priceFloat),
				Currency: "USD",
			},
			Modifiers: Itemmodifiers,
		})
	}

	return dto.Order{
		Order: RequestPayload,
	}

}
