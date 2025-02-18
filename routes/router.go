package routes

import (
	"squarepos/handlers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router{
	router:=mux.NewRouter()
	router.HandleFunc("/order",handlers.CreateOrder).Methods("POST")
	return router;
}