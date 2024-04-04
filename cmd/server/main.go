package main

import (
	"fmt"
	"net/http"
	"os"

	"rate-exchange-api/internal/currencyexchange"
	"rate-exchange-api/internal/handler"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	svc := currencyexchange.InitializeService()

	exchangeHandler := handler.NewExchangeHandler(svc)

	http.HandleFunc("/convert", exchangeHandler)

	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
