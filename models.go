package main

type HistoricalPriceResponse struct {
	Symbol     string             `json:"symbol"`
	Historical []HistoricalRecord `json:"historical"`
}

type HistoricalRecord struct {
	Date  string  `json:"date"` // Keep it as string
	Close float64 `json:"close"`
}
