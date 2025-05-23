package main

import (
	"bufio"
	"fmt"
	"go-Converter/engine"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello! Welcome to the Converter!")
	fmt.Println("Please enter your conversion query in the format: <amount> <FROM_CURRENCY> to <TO_CURRENCY>")
	fmt.Println("You can convert between currencies like USD, EUR, RUB, GBP.")
	fmt.Println("-----------------------------------------------")
	fmt.Println("For example: 100 USD to EUR")
	fmt.Println("For example: 50 EUR to RUB")
	fmt.Println("For example: 1000 RUB to USD")
	fmt.Println("For example: 20 GBP to USD")
	fmt.Println("-----------------------------------------------")
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

		result, err := engine.ConvertCurrency(amount, fromCurrency, toCurrency)
		if err != nil {
			fmt.Printf("Conversion Error: %v\n", err)
			continue
		}

		// Clarified output: "Result is: <converted_amount> <TARGET_CURRENCY>"
		fmt.Printf("Result is: %.2f %s\n", result, toCurrency)
	}
}
