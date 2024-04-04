package handler

import (
	"encoding/json"
	"net/http"
	"regexp"

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

		if source == "" || target == "" {
			http.Error(w, "Source and target currencies must be specified", http.StatusBadRequest)
			return
		}

		matched, err := regexp.MatchString(`^(\d+|\d{1,3}(,\d{3})*)(\.\d+)?$`, amount)
		if err != nil || !matched {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

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
