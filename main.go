package main

import (
	"fmt"
	"log"
	"net/http"
	"squarepos/routes"
)

func main() {

	router := routes.GetRoutes()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
	// Check why the above function is blocking the
	fmt.Println("server running on port 8080")
}
