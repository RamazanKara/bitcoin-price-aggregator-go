package bitcoin_price_prognosis

import (
	"errors"
	"fmt"
	"math"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
)

// Update the function signature to accept a slice of float64
func PrognoseNextDayPrice(prices []float64, config aggregator.Config) (string, string, error) {
	n := len(prices)
	if n < 30 { // Adjusted to check the length of preprocessed data
		return "0.00", "", errors.New("insufficient data points for accurate prognosis")
	}

	var sumX, sumY, sumXY, sumXX float64
	for i, value := range prices {
		x := float64(i + 1)
		sumX += x
		sumY += value
		sumXY += x * value
		sumXX += x * x
	}

	// Linear regression calculations
	slope := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXX - sumX*sumX)
	intercept := (sumY - slope*sumX) / float64(n)

	// Predict next day price
	nextDayPrice := slope*float64(n+1) + intercept

	// Trend determination
	trend := determineTrend(slope)

	formattedNextDayPrice := fmt.Sprintf("%.2f %s", nextDayPrice, config.VsCurrency)
	return formattedNextDayPrice, trend, nil
}

func determineTrend(slope float64) string {
	if math.Abs(slope) > 0.1 { // Threshold for significant trend
		if slope > 0 {
			return "upwards"
		} else {
			return "downwards"
		}
	}
	return "stable"
}
