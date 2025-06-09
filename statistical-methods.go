package main

import (
	"math"
	"sort"
)

func (app *application) CalculateStockCentralTendencies(data []HistoricalRecord) *MeasuresOfCentralTendencies {

	var mean float64
	var sum float64
	var median float64
	var variance float64
	count := len(data)
	prices := []float64{}

	// Mean
	for _, dataPoint := range data {
		sum += dataPoint.Close
	}
	mean = sum / float64(count)

	// Median
	sort.Float64s(prices)
	if len(prices)%2 == 0 {
		mid1 := prices[count/2]
		mid2 := prices[(count/2)-1]
		median = (mid1 + mid2) / 2
	} else {
		median = prices[count/2]
	}

	// Variance
	for _, dataPoint := range data {
		deviation := dataPoint.Close - mean
		variance += deviation * deviation
	}
	variance /= float64(count - 1)

	// standard deviation
	stdDev := math.Sqrt(variance)

	CentralTendencies := MeasuresOfCentralTendencies{
		Mean:              mean,
		Median:            median,
		Count:             count,
		Variance:          variance,
		StandardDeviation: stdDev,
	}

	return &CentralTendencies
}
