package routes

import (
	"encoding/json"
	"net/http"
	"squarepos/auth"
	"squarepos/handlers"
	"squarepos/middleware"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := handlers.CreateOrder(writer, request); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}))).Methods(http.MethodPost)
	//router.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetOrderById))).Methods(http.MethodGet)
	router.Handle("/orders/{id}", middleware.AuthMiddleware(http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			order, err := handlers.GetOrderById(writer, request)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			json.NewEncoder(writer).Encode(order)
		}))).Methods(http.MethodGet)

	router.Handle("/orders/payments", middleware.AuthMiddleware(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if err := handlers.MakePayment(writer, request); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	}))).Methods(http.MethodPost)
	router.HandleFunc("/register", auth.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login", auth.Login).Methods(http.MethodPost)
	return router
}
