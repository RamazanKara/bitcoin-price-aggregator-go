package price_preprocessing

import (
	"math"
	"sort"
	"strconv"

	aggregator "github.com/RamazanKara/bitcoin-price-aggregator-go/pkg/price-aggregator"
)

func PreprocessData(prices []aggregator.BitcoinPrice) ([]float64, error) {
	var processedPrices []float64

	// Convert prices to float and handle any missing values
	for _, price := range prices {
		if price.Value == "" {
			continue // Skipping missing values
		}
		value, err := strconv.ParseFloat(price.Value, 64)
		if err != nil {
			return nil, err
		}
		processedPrices = append(processedPrices, value)
	}

	// Remove outliers
	processedPrices = removeOutliers(processedPrices)
	return processedPrices, nil
}

func removeOutliers(data []float64) []float64 {
	mean, std := calculateMeanAndStd(data)
	var filteredData []float64
	for _, value := range data {
		if math.Abs(value-mean) < 2*std { // Keeping values within 2 standard deviations
			filteredData = append(filteredData, value)
		}
	}
	return filteredData
}

func calculateMeanAndStd(data []float64) (mean, std float64) {
	total := 0.0
	for _, value := range data {
		total += value
	}
	mean = total / float64(len(data))

	var varianceSum float64
	for _, value := range data {
		varianceSum += math.Pow(value-mean, 2)
	}
	std = math.Sqrt(varianceSum / float64(len(data)-1))
	return
}

func min(data []float64) float64 {
	sort.Float64s(data)
	return data[0]
}

func max(data []float64) float64 {
	sort.Float64s(data)
	return data[len(data)-1]
}
