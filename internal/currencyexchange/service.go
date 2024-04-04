package currencyexchange

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"rate-exchange-api/internal/model"
)

type CurrencyExchangeService interface {
	ConvertCurrency(source string, target string, amountStr string) (string, error)
}

type currencyExchangeServiceImpl struct {
	data model.ExchangeRateMap
}

func (c currencyExchangeServiceImpl) ConvertCurrency(source string, target string, amountStr string) (string, error) {
	amountStr = strings.Replace(amountStr, ",", "", -1)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return "", errors.New("invalid amount format")
	}

	rate, ok := c.data[source][target]
	if !ok {
		return "", errors.New("currency conversion rate not found")
	}

	convertedAmount := amount * rate

	roundedAmount := math.Round(convertedAmount*100) / 100
	formattedAmount := formatWithComma(roundedAmount)

	return formattedAmount, nil
}

func formatWithComma(amount float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%.2f", amount)
}

func NewCurrencyExchangeService(data model.ExchangeRateMap) CurrencyExchangeService {
	return &currencyExchangeServiceImpl{data: data}
}

func LoadExchangeRateData(filePath string) model.ExchangeRateMap {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to load exchange rate data from %s, using default. Error: %v", filePath, err)
		return model.ExchangeRateData
	}

	var rates model.ExchangeRateMap
	if err := json.Unmarshal(data, &rates); err != nil {
		log.Printf("Failed to parse exchange rate data, using default. Error: %v", err)
		return model.ExchangeRateData
	}

	return rates
}
