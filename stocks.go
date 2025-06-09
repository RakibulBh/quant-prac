package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) LastYearAverage(stock string, currency string) {

	stockYearAverageData, err := app.FetchOneYearStockData(stock)
	if err != nil {
		fmt.Printf("Error fetching one year stock data: %s\n", err.Error())
		return
	}

	centralTendencies := app.CalculateStockCentralTendencies(stockYearAverageData.Historical)

	if centralTendencies.Count > 0 {
		fmt.Printf("The mean closing price for %s in the last %d days is $%.2f\n", stock, centralTendencies.Count, centralTendencies.Mean)
		fmt.Printf("The median closing price for %s in the last %d days is $%.2f\n", stock, centralTendencies.Count, centralTendencies.Median)
		fmt.Printf("The standard deviation of closing prices for %s in the last %d days is $%.2f\n", stock, centralTendencies.Count, centralTendencies.StandardDeviation)
	} else {
		fmt.Println("No historical data to calculate average.")
	}
}

func (app *application) FetchOneYearStockData(stock string) (*HistoricalPriceResponse, error) {
	url := fmt.Sprintf("https://financialmodelingprep.com/api/v3/historical-price-full/%s?timeseries=365&serietype=line&apikey=%s", stock, app.cfg.accessKey)

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

	if len(stockYearAverageData.Historical) > 1 {
		fmt.Printf("the first date of the range is: %s\n", stockYearAverageData.Historical[0].Date)
		fmt.Printf("the last date of the range is: %s\n", stockYearAverageData.Historical[len(stockYearAverageData.Historical)-1].Date)
	}

	return &stockYearAverageData, nil
}
