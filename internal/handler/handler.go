package handler

import (
	"encoding/json"
	"net/http"

	"rate-exchange-api/internal/currencyexchange"
)

type ExchangeResponse struct {
	Msg    string `json:"msg"`
	Amount string `json:"amount"`
}

func NewExchangeHandler(svc currencyexchange.CurrencyExchangeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.URL.Query().Get("source")
		target := r.URL.Query().Get("target")
		amount := r.URL.Query().Get("amount")

		convertedAmount, err := svc.ConvertCurrency(source, target, amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ExchangeResponse{
			Msg:    "success",
			Amount: convertedAmount,
		}
		w.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
