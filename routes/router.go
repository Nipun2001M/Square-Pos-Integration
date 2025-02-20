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
	router.Handle("/orders",middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateOrder))).Methods(http.MethodPost)
	router.Handle("/orders/{id}",middleware.AuthMiddleware(http.HandlerFunc(handlers.GetOrderById))).Methods(http.MethodGet)
	router.Handle("/orders/payments",middleware.AuthMiddleware(http.HandlerFunc(handlers.MakePayment))).Methods(http.MethodPost)
	router.HandleFunc("/register",auth.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login",auth.Login).Methods(http.MethodPost)
	return router;
}