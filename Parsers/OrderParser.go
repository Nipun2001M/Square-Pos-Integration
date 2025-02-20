package parsers

import (
	"squarepos/dto"
	"strconv"
)

func OrderParser(IncomingRes dto.OrderResponse) dto.OrderResponseOut{
	var orderRes dto.OrderResponseOut
	orderRes.Id=IncomingRes.Order.ID
	orderRes.OpenedAt=IncomingRes.Order.Created
	if IncomingRes.Order.State== "COMPLETED" {
		orderRes.Isclosed = true
	} else {
		orderRes.Isclosed = false
	}
	orderRes.Table=IncomingRes.Order.LocationId
	for _,item:=range IncomingRes.Order.LineItems{
		itemF:=dto.ItemFormated{
			Name: item.Name,
			Comment: "",
			UnitPrice: int(item.BasePrice.Amount),
			Quantity: stringToInt(item.Quantity),
			Amout: int(item.TotalMoney.Amount),
			Discounts:dto.Discount{
				Name: "Discount1",
				IsPercentage: false,
				Value: int(item.TotalDiscount.Amount),
				Amount: int(item.TotalDiscount.Amount),
			},
			
			
		}
		for _,modifier:=range item.Modifiers{
			var modifiernew=dto.Modifiers{
				Name: modifier.Name,
				UnitPrice: int(modifier.BasePrice.Amount),
				Quantity: modifier.Quantity,
				Amount:int(modifier.Amount.Amount) ,
			}
			itemF.Modifiers=append(itemF.Modifiers, modifiernew)
		}
		orderRes.Items=append(orderRes.Items, itemF)

		
	}

	orderRes.Totals=dto.Totals{
		Discounts: int(IncomingRes.Order.NetAmout.Discount.Amount),
		Due:int(IncomingRes.Order.Due.Amount),
		Tax:int(IncomingRes.Order.Tax.Amount),
		ServiceCharge:int(IncomingRes.Order.NetAmout.ServiceCharge.Amount),
		Tips:int(IncomingRes.Order.NetAmout.Tips.Amount),
		Total:int(IncomingRes.Order.TotalMoney.Amount),
		}
	return orderRes

}

func stringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}