//go:build wireinject

package currencyexchange

import (
	"github.com/google/wire"
)

func InitializeService() CurrencyExchangeService {
	wire.Build(NewCurrencyExchangeService, NewExchangeRateData)
	return nil
}
