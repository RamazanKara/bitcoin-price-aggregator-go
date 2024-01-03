package bitcoin_price_prognosis

import (
	"errors"
	"math"
	"github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/bitcoin_price_aggregator"
	"strconv"
	"fmt"
)

func PrognoseNextDayPrice(prices []bitcoin_price_aggregator.BitcoinPrice, config bitcoin_price_aggregator.Config) (string, string, error) {
	n := len(prices)
	if n == 0 {
		// Return formatted string with "0.00" and the currency
		return fmt.Sprintf("0.00 %s", config.VsCurrency), "", errors.New("no prices available for prognosis")
	}

	var sumX, sumY, sumXY, sumXX float64
	for i, price := range prices {
		x := float64(i + 1)

		y, err := strconv.ParseFloat(price.Value, 64)
		if err != nil {
			return "", "", fmt.Errorf("invalid price value: %v", err)
		}
		sumX += x
		sumY += y
		sumXY += x * y
		sumXX += x * x
	}

	slope := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXX - sumX*sumX)
	intercept := (sumY - slope*sumX) / float64(n)

	nextDayPrice := slope*float64(n+1) + intercept
	trend := "stable"
	if math.Abs(slope) > 0.01 { // Threshold for significant trend
		if slope > 0 {
			trend = "upwards"
		} else {
			trend = "downwards"
		}
	}

	formattedNextDayPrice := fmt.Sprintf("%.2f %s", nextDayPrice, config.VsCurrency)

	return formattedNextDayPrice, trend, nil
}
