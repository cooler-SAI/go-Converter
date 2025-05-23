package main

import (
	"fmt"
	"strings"
)

var hardcodedRates = map[string]float64{
	"USD_EUR": 0.92, // 1 USD = 0.92 EUR
	"EUR_USD": 1.08, // 1 EUR = 1.08 USD
	"USD_RUB": 90.0, // 1 USD = 90.0 RUB
	"RUB_USD": 1.0 / 90.0,
	"EUR_RUB": 98.0, // 1 EUR = 98.0 RUB
	"RUB_EUR": 1.0 / 98.0,
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Hardcoded
	rateKey := fmt.Sprintf("%s_%s", fromCurrency, toCurrency)
	if rate, ok := hardcodedRates[rateKey]; ok {
		return amount * rate, nil
	}
	return 0, fmt.Errorf("can't convert %s to %s: cource doesn't found", fromCurrency, toCurrency)
}

func main() {
	fmt.Println("Hello! Please enter data in format: " +
		"")
}
