package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PaymentRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)

	// simulate payment processing
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment successful"))
}

func main() {
	http.HandleFunc("/pay", paymentHandler)
	fmt.Println("Payment Service is running on port 8081")
	http.ListenAndServe(":8082", nil)
}
