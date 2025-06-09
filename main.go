package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	accessKey string
}

type application struct {
	cfg *config
}

func main() {
	godotenv.Load()

	cfg := &config{
		accessKey: os.Getenv("API_KEY"),
	}

	if cfg.accessKey == "" {
		fmt.Print(cfg.accessKey)
		return
	}

	app := &application{
		cfg: cfg,
	}

	// app.LastYearAverage("AAPL", "GBP")
	// app.SimulateDiceRolls(10000)
	app.CalculateSMA(30, "AAPL")
}
