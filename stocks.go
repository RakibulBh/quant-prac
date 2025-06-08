package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
)

func (app *application) LastYearAverage(stock string, currency string) {

	stockYearAverageData, err := app.FetchOneYearStockData(stock)
	if err != nil {
		fmt.Printf("Error fetching one year stock data: %s\n", err.Error())
	}

	// data
	sum := 0.0
	count := len(stockYearAverageData.Historical)
	median := 0.0
	var mean float64
	var variance float64
	var sd float64

	// Mean
	prices := []float64{}
	for _, data := range stockYearAverageData.Historical {
		prices = append(prices, data.Close)
		sum += data.Close
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
	for _, v := range prices {
		variance += (v - mean) * (v - mean)
	}
	variance /= (float64(count) - 1)

	// Standard deviation
	sd = math.Sqrt(variance)

	if count > 0 {
		fmt.Printf("The mean closing price for %s in the last %d days is $%.2f\n", stock, count, sum/float64(count))
		fmt.Printf("The median closing price for %s in the last %d days is $%.2f\n", stock, count, median)
		fmt.Printf("The standard deviation of closing prices for %s in the last %d days is $%.2f\n", stock, count, sd)
	} else {
		fmt.Println("No historical data to calculate average.")
	}
}

func (app *application) FetchOneYearStockData(stock string) (*HistoricalPriceResponse, error) {
	url := fmt.Sprintf("https://financialmodelingprep.com/api/v3/historical-price-full/%s?timeseries=365&serietype=line&apikey=jgU0TAp8KhQyiM0pnsZzycqa0wI6L5Ky", stock)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	defer res.Body.Close()

	var stockYearAverageData HistoricalPriceResponse
	if err := app.ReadJSON(resBody, &stockYearAverageData); err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return nil, err
	}

	return &stockYearAverageData, nil
}
