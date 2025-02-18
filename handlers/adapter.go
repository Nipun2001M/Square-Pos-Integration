package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"squarepos/apicall"
	"squarepos/dto"
	"squarepos/response"
)

func CreateOrder(w http.ResponseWriter, req *http.Request) {
	var OrderReq dto.Order
	err := json.NewDecoder(req.Body).Decode(&OrderReq)
	if err != nil {
		http.Error(w, "Error in decoding req body", http.StatusBadRequest)
	}

	client := apicall.GetClient()
	data, error := client.ApiCall(http.MethodPost, "orders",OrderReq)

	if error != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	var res response.OrderResponse
	fmt.Println(res)
	json.Unmarshal(data,&res)
	json.NewEncoder(w).Encode(res)

}
