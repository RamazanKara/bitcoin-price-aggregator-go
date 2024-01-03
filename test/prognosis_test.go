package test
import (
	"testing"
	"github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/bitcoin_price_aggregator"
	"github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/bitcoin_price_prognosis"
)

func TestPrognoseNextDayPrice(t *testing.T) {
    testConfig := bitcoin_price_aggregator.ApiConfig()
	currencySuffix := " " + testConfig.VsCurrency

    testCases := []struct {
        name     string
        prices   []bitcoin_price_aggregator.BitcoinPrice
        expected string
        trend    string
    }{
        {
            name: "Empty Prices",
            prices: []bitcoin_price_aggregator.BitcoinPrice{},
            expected: "0.00" + currencySuffix,
            trend: "",
        },
        {
            name: "Upward Trend",
            prices: []bitcoin_price_aggregator.BitcoinPrice{
                {Value: "100.00"},
                {Value: "200.00"},
                {Value: "300.00"},
            },
            expected: "400.00" + currencySuffix,
			trend: "upwards",
        },
        {
            name: "Downward Trend",
            prices: []bitcoin_price_aggregator.BitcoinPrice{
                {Value: "300.00"},
                {Value: "200.00"},
                {Value: "100.00"},
            },
            expected: "0.00" + currencySuffix,
            trend: "downwards",
        },
        {
            name: "Stable Trend",
            prices: []bitcoin_price_aggregator.BitcoinPrice{
                {Value: "200.00"},
                {Value: "200.00"},
                {Value: "200.00"},
            },
            expected: "200.00" + currencySuffix,
            trend: "stable",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            prognosis, trend, err := bitcoin_price_prognosis.PrognoseNextDayPrice(tc.prices, testConfig)
            if len(tc.prices) == 0 && err == nil {
                t.Error("Expected an error for empty price data")
            }
            if len(tc.prices) > 0 && err != nil {
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