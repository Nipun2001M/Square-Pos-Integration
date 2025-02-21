package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"squarepos/Parsers"
	"squarepos/apiClient"
	"squarepos/auth"
	"squarepos/dto"
	"squarepos/middleware"
)

// error
func CreateOrder(w http.ResponseWriter, req *http.Request) error {
	var OrderReq qlub.SubmitOrderCommand
	fmt.Println("CreateOrder ", OrderReq)
	err := json.NewDecoder(req.Body).Decode(&OrderReq)
	if err != nil {
		http.Error(w, "Error in decoding req body", http.StatusBadRequest)
		return err
	}
	claims, ok := req.Context().Value(middleware.UserContextKey).(*auth.Claims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: No valid claims", http.StatusUnauthorized)
		return errors.New("Unauthorized: No valid claims")

	}
	client := apiClient.GetClient()
	_, error := client.ApiCall(http.MethodPost, "orders", OrderReq, claims.AccessToken)
	if error != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
		return errors.New("error in api call func")
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
	//formattedRes := dto.OrderResponse{}
	//json.Unmarshal(data, &formattedRes)
	//json.NewEncoder(w).Encode(parsers.OrderParser(formattedRes))
	//json.NewEncoder(w).Encode(formattedRes)

}

// (order qlub.Order, err error)
func GetOrderById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	claims, ok := req.Context().Value(middleware.UserContextKey).(*auth.Claims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: No valid claims", http.StatusUnauthorized)
		return
	}
	client := apiClient.GetClient()
	data, error := client.ApiCall(http.MethodGet, fmt.Sprintf("orders/%s", id), nil, claims.AccessToken)
	if error != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	formattedRes := dto.OrderResponse{}
	json.Unmarshal(data, &formattedRes)
	json.NewEncoder(w).Encode(parsers.OrderParser(formattedRes))

}

func MakePayment(w http.ResponseWriter, req *http.Request) {
	//var PaymentReq qlub.UpdatePaymentStatusCommand
	var PaymentReq dto.PaymentRequest
	err := json.NewDecoder(req.Body).Decode(&PaymentReq)
	PaymentReq.IdempotencyKey = uuid.New().String()
	claims, ok := req.Context().Value(middleware.UserContextKey).(*auth.Claims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: No valid claims", http.StatusUnauthorized)
		return

	}
	if err != nil {
		http.Error(w, "error occured in decoding payment body", http.StatusBadRequest)
		return
	}
	client := apiClient.GetClient()
	data, err := client.ApiCall(http.MethodPost, "payments", &PaymentReq, claims.AccessToken)
	if err != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var defaultRes map[string]interface{}
	json.Unmarshal(data, &defaultRes)
	json.NewEncoder(w).Encode(defaultRes)

}
