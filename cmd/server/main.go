package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"rate-exchange-api/internal/currencyexchange"
	"rate-exchange-api/internal/handler"
)

func main() {
	rootDir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal("Failed to determine project root directory:", err)
	}

	err = godotenv.Load(filepath.Join(rootDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

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
