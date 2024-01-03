package main

import (
	"fmt"
	"log"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
	prognosis "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-prognosis"
)

func main() {
	config := aggregator.ApiConfig()
	prices, err := aggregator.FetchBitcoinPrices(config)
	if err != nil {
		log.Fatalf("Error fetching Bitcoin prices: %v", err)
	}

	for _, price := range prices {
		fmt.Printf("Date: %s, Price: %s %s\n", price.Date, price.Value, config.VsCurrency)
	}

	prognosis, trend, err := prognosis.PrognoseNextDayPrice(prices, config)
	if err != nil {
		log.Fatalf("Error calculating prognosis: %v", err)
	}

	fmt.Printf("Prognosis for the next day: %s (%s trend)\n", prognosis, trend)
}
