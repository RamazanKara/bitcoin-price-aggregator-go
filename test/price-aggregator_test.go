package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
)

func TestFetchBitcoinPrices(t *testing.T) {
	// Mock server to simulate API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := `{"prices":[[1609459200000, 29031.34], [1609545600000, 29374.15], [1609632000000, 32127.27]]}`
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Create a local config for the test
	testConfig := aggregator.Config{
		CoinGeckoAPIURL: mockServer.URL,
		VsCurrency:      "eur",
	}

	prices, err := aggregator.FetchBitcoinPrices(testConfig)
	if err != nil {
		t.Fatalf("Failed to fetch Bitcoin prices: %v", err)
	}

	if len(prices) != 3 {
		t.Errorf("Expected 3 prices, got %d", len(prices))
	}
}
