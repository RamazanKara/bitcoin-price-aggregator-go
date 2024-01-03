package main

import (
	"fmt"
	"github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/bitcoin_price_aggregator"
	"github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/bitcoin_price_prognosis"
	"log"
)

func main() {
	config := bitcoin_price_aggregator.ApiConfig()
	prices, err := bitcoin_price_aggregator.FetchBitcoinPrices(config)
	if err != nil {
		log.Fatalf("Error fetching Bitcoin prices: %v", err)
	}

	for _, price := range prices {
		fmt.Printf("Date: %s, Price: %s %s\n", price.Date, price.Value, config.VsCurrency)
	}

	prognosis, trend, err := bitcoin_price_prognosis.PrognoseNextDayPrice(prices, config)
	if err != nil {
		log.Fatalf("Error calculating prognosis: %v", err)
	}

	fmt.Printf("Prognosis for the next day: %s (%s trend)\n", prognosis, trend)
}
