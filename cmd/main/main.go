package main

import (
	"fmt"
	"log"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
	preprocessing "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-preprocessing"
	prognosis "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-prognosis"
)

func main() {
	config := aggregator.ApiConfig()

	// Fetch Bitcoin prices
	prices, err := aggregator.FetchBitcoinPrices(config)
	if err != nil {
		log.Fatalf("Error fetching Bitcoin prices: %v", err)
	}

	// Display fetched prices
	for _, price := range prices {
		fmt.Printf("Date: %s, Price: %s %s\n", price.Date, price.Value, config.VsCurrency)
	}

	// Preprocess the prices
	preprocessedPrices, err := preprocessing.PreprocessData(prices)
	if err != nil {
		log.Fatalf("Error in preprocessing data: %v", err)
	}

	// Calculate and display the prognosis
	nextDayPricePrognosis, trend, err := prognosis.PrognoseNextDayPrice(preprocessedPrices, config)
	if err != nil {
		log.Fatalf("Error calculating prognosis: %v", err)
	}
	fmt.Printf("Prognosis for the next day: %s (%s trend)\n", nextDayPricePrognosis, trend)
}
