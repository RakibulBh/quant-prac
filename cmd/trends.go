package main

import (
	"fmt"
	"math"
)

func (app *application) CalculateSMA(window int, stock string) float64 {

	stocks, err := app.FetchOneYearStockData(stock)
	if err != nil {
		fmt.Printf("Error fetching one year stock data: %s\n", err.Error())
		return 0
	}

	history := stocks.Historical
	dataLen := len(history)

	if dataLen < window {
		fmt.Printf("Not enough data to calculate %d-day SMA\n", window)
		return 0
	}

	average := 0.0
	for i := range window {
		average += history[i].Close
	}

	sma := math.Round((average/float64(window))*100) / 100

	fmt.Printf("The Simple Moving Average of %s in the last %d days is $%.2f\n", stock, window, sma)

	return sma
}

// TODO
// func (app *application) ComputeAutoCorrelationLag1(data HistoricalPriceResponse) {

// 	centralTendencies := app.CalculateStockCentralTendencies(data.Historical)

// }
