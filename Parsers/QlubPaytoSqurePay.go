package parsers

import (
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	qkit "github.com/clubpay/qlubkit-go"
	"github.com/google/uuid"
	"squarepos/dto"
)

func QlubPaytoSquarePay(Incoming qlub.UpdatePaymentStatusCommand) dto.PaymentRequest {
	var PaymentReq dto.PaymentRequest
	PaymentReq.IdempotencyKey = uuid.New().String()
	PaymentReq.AmountMoney = dto.Money{
		Amount:   qkit.StrToFloat64(Incoming.TotalAmount),
		Currency: "USD",
	}
	PaymentReq.LocationID = Incoming.TransactionID
	PaymentReq.SourceId = "ccof:customer-card-id-ok"
	PaymentReq.OrderId = Incoming.OrderID

	return PaymentReq

}
