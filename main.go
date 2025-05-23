package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func convertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
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

func main() {
	fmt.Println("Hello! Welcome to the Converter!")
	fmt.Println("Please enter your conversion query in the format: <amount> <FROM_CURRENCY> to <TO_CURRENCY>")
	fmt.Println("For example: 100 EUR to RUB")
	fmt.Println("Type 'exit' or 'quit' to close the application.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter query: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			break // Exit if scanner fails (e.g., EOF)
		}

		userInput := strings.TrimSpace(scanner.Text())
		if strings.ToLower(userInput) == "exit" || strings.ToLower(userInput) == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		parts := strings.Fields(userInput) // Splits by whitespace

		// Expected format: "100", "USD", "to", "EUR" (4 parts)
		if len(parts) != 4 || strings.ToLower(parts[2]) != "to" {
			fmt.Println("Invalid format. Please use: <amount> <FROM_CURRENCY> to <TO_CURRENCY>")
			fmt.Println("For example: 100 EUR to RUB")
			continue
		}

		amountStr := parts[0]
		fromCurrency := strings.ToUpper(parts[1])
		toCurrency := strings.ToUpper(parts[3])

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Printf("Error: Amount '%s' is not a valid number.\n", amountStr)
			continue
		}

		if amount < 0 {
			fmt.Println("Error: Amount cannot be negative.")
			continue
		}

		result, err := convertCurrency(amount, fromCurrency, toCurrency)
		if err != nil {
			fmt.Printf("Conversion Error: %v\n", err)
			continue
		}

		// Clarified output: "Result is: <converted_amount> <TARGET_CURRENCY>"
		fmt.Printf("Result is: %.2f %s\n", result, toCurrency)
	}
}
