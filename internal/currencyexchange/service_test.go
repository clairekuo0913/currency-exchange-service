package currencyexchange

import (
	"testing"

	"github.com/stretchr/testify/require"
	"rate-exchange-api/internal/model"
)

func TestCurrencyExchangeService_ConvertCurrency(t *testing.T) {
	r := require.New(t)

	data := model.ExchangeRateData
	service := NewCurrencyExchangeService(data)

	tests := []struct {
		name         string
		source       string
		target       string
		amountStr    string
		expected     string
		expectingErr bool
	}{
		{
			name:         "TWD to JPY",
			source:       model.CurrencyCodeTWD,
			target:       model.CurrencyCodeJPY,
			amountStr:    "100",
			expected:     "366.90", // TWD to JPY is 3.669
			expectingErr: false,
		},
		{
			name:         "USD to TWD with comma",
			source:       model.CurrencyCodeUSD,
			target:       model.CurrencyCodeTWD,
			amountStr:    "1,000",
			expected:     "30,444.00", // USD to TWD is 30.444
			expectingErr: false,
		},
		{
			name:         "USD to JPY without comma",
			source:       model.CurrencyCodeUSD,
			target:       model.CurrencyCodeJPY,
			amountStr:    "1525",
			expected:     "170,496.53", // USD to JPY is 111.801
			expectingErr: false,
		},
		{
			name:         "Invalid amount format",
			source:       model.CurrencyCodeUSD,
			target:       model.CurrencyCodeJPY,
			amountStr:    "invalid",
			expected:     "",
			expectingErr: true,
		},
		{
			name:         "Conversion rate not found",
			source:       "ABC",
			target:       "DEF",
			amountStr:    "100",
			expected:     "",
			expectingErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ConvertCurrency(tt.source, tt.target, tt.amountStr)

			if tt.expectingErr {
				r.Error(err)
			} else {
				r.NoError(err)
				r.Equal(tt.expected, result)
			}
		})
	}
}
