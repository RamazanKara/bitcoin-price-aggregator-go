package test

import (
	"testing"
	"fmt"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
	preprocessing "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-preprocessing"
	prognosis "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-prognosis"
)

func generateSequentialPrices(num int, start, increment float64) []aggregator.BitcoinPrice {
	var prices []aggregator.BitcoinPrice
	for i := 0; i < num; i++ {
		price := start + float64(i)*increment
		prices = append(prices, aggregator.BitcoinPrice{Value: fmt.Sprintf("%.2f", price)})
	}
	return prices
}

func generateStablePrices(num int, value float64) []aggregator.BitcoinPrice {
	var prices []aggregator.BitcoinPrice
	for i := 0; i < num; i++ {
		prices = append(prices, aggregator.BitcoinPrice{Value: fmt.Sprintf("%.2f", value)})
	}
	return prices
}

func TestPrognoseNextDayPrice(t *testing.T) {
	testConfig := aggregator.ApiConfig()
	currencySuffix := " " + testConfig.VsCurrency

	testCases := []struct {
		name     string
		prices   []aggregator.BitcoinPrice
		expected string
		trend    string
	}{
		{
			name:     "Empty Prices",
			prices:   []aggregator.BitcoinPrice{},
			expected: "0.00",
			trend:    "",
		},
		{
			name: "Upward Trend",
			prices: generateSequentialPrices(30, 100.00, 10.00),
			expected: "400.00" + currencySuffix,
			trend:    "upwards",
		},
		{
			name: "Downward Trend",
			prices: generateSequentialPrices(30, 400.00, -10.00),
			expected: "100.00" + currencySuffix,
			trend:    "downwards",
		},
		{
			name: "Stable Trend",
			prices: generateSequentialPrices(30, 200.00, 0.001),
			expected: "200.03" + currencySuffix,
			trend:    "stable",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			preprocessedPrices, err := preprocessing.PreprocessData(tc.prices)
			if err != nil {
				t.Fatalf("Preprocessing failed: %v", err)
			}

			prognosis, trend, err := prognosis.PrognoseNextDayPrice(preprocessedPrices, testConfig)
			if len(preprocessedPrices) == 0 && err == nil {
				t.Error("Expected an error for empty or invalid preprocessed data")
			}
			if len(preprocessedPrices) > 0 && err != nil {
				t.Errorf("Did not expect an error: %v", err)
			}
			if prognosis != tc.expected {
				t.Errorf("Expected prognosis %s, got %s", tc.expected, prognosis)
			}
			if trend != tc.trend {
				t.Errorf("Expected trend %s, got %s", tc.trend, trend)
			}
		})
	}
}
