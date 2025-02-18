package main

import (
	"fmt"
	"log"
	"net/http"
	"squarepos/routes"
)

func main() {

	router := routes.GetRoutes()
	fmt.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}