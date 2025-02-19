package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"squarepos/apicall"
	"squarepos/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, req *http.Request) {
	var OrderReq dto.Order
	err := json.NewDecoder(req.Body).Decode(&OrderReq)
	if err != nil {
		http.Error(w, "Error in decoding req body", http.StatusBadRequest)
	}

	client := apicall.GetClient()
	data, error := client.ApiCall(http.MethodPost, "orders", OrderReq)

	if error != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	var res dto.OrderResponse
	var defaultRes map[string]interface{}
	// Have a dto here
	fmt.Println(res)
	json.Unmarshal(data,&defaultRes)
	json.NewEncoder(w).Encode(defaultRes)
}

func GetOrderById(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req)
	id := params["id"] 
	client:=apicall.GetClient()
	data,error:=client.ApiCall(http.MethodGet,fmt.Sprintf("orders/%s",id),nil)
	if error != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	var res dto.OrderResponse
	var defaultRes map[string]interface{}
	fmt.Println(res)
	json.Unmarshal(data,&defaultRes)
	json.NewEncoder(w).Encode(defaultRes)



}

func MakePayment(w http.ResponseWriter,req * http.Request){
	var PaymentReq dto.PaymentRequest
	err:=json.NewDecoder(req.Body).Decode(&PaymentReq)
	PaymentReq.IdempotencyKey=uuid.New().String()
	if err!=nil{
		http.Error(w,"error occured in decoding payment body",http.StatusBadRequest)
		return
	}
	client:=apicall.GetClient()
	data,err:=client.ApiCall(http.MethodPost,"payments",&PaymentReq)
	if err!=nil{
		http.Error(w,"error in api call func",http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var res dto.OrderResponse
	var defaultRes map[string]interface{}
	fmt.Println(res)
	json.Unmarshal(data,&defaultRes)
	json.NewEncoder(w).Encode(defaultRes)




}
