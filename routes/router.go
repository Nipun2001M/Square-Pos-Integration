package routes

import (
	"net/http"
	"squarepos/handlers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router{
	router:=mux.NewRouter()
	//create order
	router.HandleFunc("/orders",handlers.CreateOrder).Methods(http.MethodPost)
	//get order by order id
	router.HandleFunc("/orders/{id}",handlers.GetOrderById).Methods(http.MethodGet)
	//make payment
	router.HandleFunc("/orders/payments",handlers.MakePayment).Methods(http.MethodPost)
	return router;
}