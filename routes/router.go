package routes

import (
	"net/http"
	"squarepos/handlers"

	"github.com/gorilla/mux"
)

// Try to handle this only using standard libraries
func GetRoutes() *mux.Router{
	router:=mux.NewRouter()
	router.HandleFunc("/orders",handlers.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/orders/{id}",handlers.GetOrderById).Methods(http.MethodGet)
	router.HandleFunc("/orders/payments",handlers.MakePayment).Methods(http.MethodPost)
	return router;
}