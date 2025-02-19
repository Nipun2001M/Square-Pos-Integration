package main

import (
	"fmt"
	"log"
	"net/http"
	"squarepos/database"
	"squarepos/routes"
)

func main() {

	database.Connect()
	router := routes.GetRoutes()
	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("server running on port 8080")

	select {}
	// Check why the above function is blocking the
}
