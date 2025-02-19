package routes

import (
	"net/http"
	"squarepos/auth"
	"squarepos/handlers"
	"squarepos/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router{
	router:=mux.NewRouter()
	router.HandleFunc("/orders",handlers.CreateOrder).Methods(http.MethodPost)
	router.Handle("/orders/{id}",middleware.AuthMiddleware(http.HandlerFunc(handlers.GetOrderById))).Methods(http.MethodGet)
	router.HandleFunc("/orders/payments",handlers.MakePayment).Methods(http.MethodPost)
	router.HandleFunc("/register",auth.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login",auth.Login).Methods(http.MethodPost)
	return router;
}