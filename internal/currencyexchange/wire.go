//go:build wireinject

package currencyexchange

import (
	"os"

	"github.com/google/wire"
	"rate-exchange-api/internal/model"
)

func ProvideExchangeRateData() model.ExchangeRateMap {
	filePath := os.Getenv("RATE_EXCHANGE_JSON_PATH")
	return LoadExchangeRateData(filePath)
}

func InitializeService() CurrencyExchangeService {
	wire.Build(NewCurrencyExchangeService, ProvideExchangeRateData)
	return nil
}
