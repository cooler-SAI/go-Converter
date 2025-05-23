package engine

import (
	"fmt"
	"strings"
)

// hardcodedRates stores the predefined exchange rates.
// It's kept private to the engine package (starts with a lowercase letter).
// If you needed to modify rates from outside, you might add functions to do so.
var hardcodedRates = map[string]float64{
	"USD_EUR": 0.92,
	"EUR_USD": 1.08,
	"USD_RUB": 90.0,
	"RUB_USD": 1.0 / 90.0,
	"EUR_RUB": 98.0,
	"RUB_EUR": 1.0 / 98.0,
	"USD_GBP": 0.79,
	"GBP_USD": 1.27,
	// Add more currencies as needed
}

// ConvertCurrency performs the currency conversion.
// This function is exported (starts with an uppercase letter) so it can be used by main.go.
func ConvertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	// Convert currency codes to uppercase for consistent lookup
	fromCurrency = strings.ToUpper(fromCurrency)
	toCurrency = strings.ToUpper(toCurrency)

	// If currencies are the same, no conversion needed
	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Construct the rate key for direct lookup
	rateKey := fmt.Sprintf("%s_%s", fromCurrency, toCurrency)
	if rate, ok := hardcodedRates[rateKey]; ok {
		return amount * rate, nil
	}

	// Attempt to find an inverse rate if a direct rate is not found. For example, if
	// user wants EUR_USD but we only have USD_EUR. Note: The current hardcodedRates
	// map explicitly includes inverse rates, so this block might seem redundant for
	// the current data, but it's good practice if the rate map wasn't guaranteed to
	// have all inverses.
	inverseRateKey := fmt.Sprintf("%s_%s", toCurrency, fromCurrency)
	if inverseRate, ok := hardcodedRates[inverseRateKey]; ok && inverseRate != 0 {
		return amount * (1.0 / inverseRate), nil
	}

	// If no direct or inverse rate is found, return an error.
	// Future enhancements could include multi-step conversion (e.g., RUB -> USD -> EUR).
	return 0, fmt.Errorf("unable to convert %s to %s: exchange rate not found or path not implemented", fromCurrency, toCurrency)
}
