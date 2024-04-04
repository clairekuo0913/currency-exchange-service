// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package currencyexchange

// Injectors from wire.go:

func InitializeService() CurrencyExchangeService {
	exchangeRateMap := NewExchangeRateData()
	currencyExchangeService := NewCurrencyExchangeService(exchangeRateMap)
	return currencyExchangeService
}