package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
	"github.com/gorilla/mux"
	"net/http"
	parsers "squarepos/Parsers"
	"squarepos/apiClient"
	"squarepos/auth"
	"squarepos/dto"
	"squarepos/middleware"
)

func CreateOrder(w http.ResponseWriter, req *http.Request) error {
	var OrderReq qlub.SubmitOrderCommand
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
	_, error := client.ApiCall(http.MethodPost, "orders", parsers.IncomingOrderToSquareParse(OrderReq), claims.AccessToken)
	if error != nil {
		return error
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func GetOrderById(w http.ResponseWriter, req *http.Request) (order qlub.Order, err error) {
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
	var defaultResponse dto.OrderResponse
	json.Unmarshal(data, &defaultResponse)
	order, err = parsers.SqureOrderToQlubOrder(defaultResponse)
	return order, err

}

func MakePayment(w http.ResponseWriter, req *http.Request) error {
	var PaymentReq qlub.UpdatePaymentStatusCommand
	err := json.NewDecoder(req.Body).Decode(&PaymentReq)
	claims, ok := req.Context().Value(middleware.UserContextKey).(*auth.Claims)
	Payload := parsers.QlubPaytoSquarePay(PaymentReq)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized: No valid claims", http.StatusUnauthorized)
		return fmt.Errorf("Unauthorized: No valid claims")

	}
	if err != nil {
		http.Error(w, "error occured in decoding payment body", http.StatusBadRequest)
		return fmt.Errorf("error occured in decoding payment body")
	}
	client := apiClient.GetClient()
	_, err = client.ApiCall(http.MethodPost, "payments", &Payload, claims.AccessToken)
	if err != nil {
		http.Error(w, "error in api call func", http.StatusBadRequest)
		return fmt.Errorf("error in api call func")
	}
	w.Header().Set("Content-Type", "application/json")
	//var defaultRes map[string]interface{}
	//json.Unmarshal(data, &defaultRes)
	//json.NewEncoder(w).Encode(defaultRes)
	return nil

}
