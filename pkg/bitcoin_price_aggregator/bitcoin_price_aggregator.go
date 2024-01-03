package bitcoin_price_aggregator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	VsCurrency      string
	CoinGeckoAPIURL string
}

func ApiConfig() Config {
	return Config{
		VsCurrency:      "eur",
		CoinGeckoAPIURL: "https://api.coingecko.com/api/v3/coins/bitcoin/market_chart/range",
	}
}

type BitcoinPrice struct {
	Date  string
	Value string
}

func FetchBitcoinPrices(config Config) ([]BitcoinPrice, error) {
	// Set the date range for the last 7 days
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -7)

	// Construct the full URL
	fullURL := fmt.Sprintf("%s?vs_currency=%s&from=%d&to=%d", config.CoinGeckoAPIURL, config.VsCurrency, startDate.Unix(), endDate.Unix())

	apiUrl, err := url.Parse(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing API URL: %v", err)
	}

	resp, err := http.Get(apiUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	prices, ok := result["prices"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}

	var bitcoinPrices []BitcoinPrice
	for _, p := range prices {
		pair, ok := p.([]interface{})
		if !ok || len(pair) != 2 {
			continue
		}

		timestamp := int64(pair[0].(float64))
		value := pair[1].(float64)
		formattedValue := fmt.Sprintf("%.2f", value)

		date := time.Unix(timestamp/1000, 0).Format("02-01-2006")

		bitcoinPrices = append(bitcoinPrices, BitcoinPrice{Date: date, Value: formattedValue})
	}

	return bitcoinPrices, nil
}
