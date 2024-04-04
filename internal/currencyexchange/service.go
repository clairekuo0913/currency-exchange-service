package currencyexchange

type CurrencyExchangeService interface {
	ConvertCurrency(source string, target string, amountStr string) (string, error)
}
