package engine

import (
	"fmt"
	"strings"
)

// (The hardcodedRates map and convertCurrency function remain the same as before)
var hardcodedRates = map[string]float64{
	"USD_EUR": 0.92,
	"EUR_USD": 1.08,
	"USD_RUB": 90.0,
	"RUB_USD": 1.0 / 90.0,
	"EUR_RUB": 98.0, // 1 EUR = 98.0 RUB (example rate)
	"RUB_EUR": 1.0 / 98.0,
	"USD_GBP": 0.79,
	"GBP_USD": 1.27,
}

func ConvertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	if fromCurrency == toCurrency {
		return amount, nil
	}

	rateKey := fmt.Sprintf("%s_%s", fromCurrency, toCurrency)
	if rate, ok := hardcodedRates[rateKey]; ok {
		return amount * rate, nil
	}

	// Try inverse rate if direct not found (e.g., if we have USD_EUR but user wants EUR_USD)
	// This is already handled by having explicit inverse rates in hardcodedRates.
	// For a more robust solution without explicit inverse rates:
	// inverseRateKey := fmt.Sprintf("%s_%s", toCurrency, fromCurrency)
	// if inverseRate, ok := hardcodedRates[inverseRateKey]; ok && inverseRate != 0 {
	// return amount * (1.0 / inverseRate), nil
	// }

	// Attempt conversion via a base currency (e.g., USD) if direct/inverse fails
	// This requires a more structured rate map (e.g., all rates to USD)
	// or a graph traversal algorithm for more complex paths.
	// For now, we'll keep it to direct rates for simplicity in this iterative step.

	return 0, fmt.Errorf("unable to convert %s to %s: exchange rate not found or path not implemented", fromCurrency, toCurrency)
}
