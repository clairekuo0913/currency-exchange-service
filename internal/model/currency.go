package model

const (
	CurrencyCodeTWD = "TWD"
	CurrencyCodeJPY = "JPY"
	CurrencyCodeUSD = "USD"
)

type ExchangeRateMap map[string]map[string]float64

var ExchangeRateData = ExchangeRateMap{
	CurrencyCodeTWD: {
		CurrencyCodeTWD: 1,
		CurrencyCodeJPY: 3.669,
		CurrencyCodeUSD: 0.03281,
	},
	CurrencyCodeJPY: {
		CurrencyCodeTWD: 0.26956,
		CurrencyCodeJPY: 1,
		CurrencyCodeUSD: 0.00885,
	},
	CurrencyCodeUSD: {
		CurrencyCodeTWD: 30.444,
		CurrencyCodeJPY: 111.801,
		CurrencyCodeUSD: 1,
	},
}
