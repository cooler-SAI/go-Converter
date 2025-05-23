# go-Converter

A simple command-line currency converter written in Go.

## Features

* Convert between different currencies.
* Interactive command-line interface.
* Uses hardcoded exchange rates (for simplicity in this example).

## Getting Started

### Prerequisites

* Go (version 1.x or later) installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### Installation

1.  **Clone the repository (if you have one set up):**
    ```bash
    git clone <your-repository-url>
    cd go-Converter
    ```
    **Alternatively, if you just have the `main.go` file:**
    Ensure `main.go` (the code we've been working on) is in your current directory.

2.  **Build the application (optional, but recommended for a distributable executable):**
    Open your terminal in the directory containing `main.go` and run:
    ```bash
    go build -o go-converter main.go
    ```
    This will create an executable file named `go-converter` (or `go-converter.exe` on Windows).

3.  **Run the application directly (without building an executable):**
    Alternatively, you can run the application directly using:
    ```bash
    go run main.go
    ```

## Usage

Once the application is running (either via the built executable `./go-converter` or `go run main.go`):

1.  You will see a welcome message:
    ```
    Hello! Welcome to the Converter!
    Please enter your conversion query in the format: <amount> <FROM_CURRENCY> to <TO_CURRENCY>
    For example: 100 EUR to RUB
    Type 'exit' or 'quit' to close the application.
    ```

2.  Enter your conversion query at the prompt:
    ```
    Enter query: <amount> <SOURCE_CURRENCY> to <TARGET_CURRENCY>
    ```
    For example:
    ```
    Enter query: 100 USD to EUR
    ```

3.  The application will display the result:
    ```
    Result is: 92.00 EUR
    ```
    (Assuming the USD to EUR rate is 0.92)

4.  To stop the application, type `exit` or `quit` and press Enter:
    ```
    Enter query: exit
    Goodbye!
    ```

## Code Overview

The application consists of a single Go file (`main.go`) with the following key components:

* **`hardcodedRates` (map):** A map storing the predefined exchange rates. The key is a string in the format `"FROMCURRENCY_TOCURRENCY"` (e.g., `"USD_EUR"`), and the value is the `float64` exchange rate.
* **`convertCurrency` function:** Takes an amount, source currency, and target currency as input. It looks up the rate in `hardcodedRates` and returns the converted amount or an error if the conversion is not possible.
* **`main` function:**
    * Prints a welcome message and instructions.
    * Enters a loop to continuously read user input.
    * Parses the input string (expected format: `<amount> <FROM_CURRENCY> to <TO_CURRENCY>`).
    * Handles "exit" or "quit" commands.
    * Calls `convertCurrency` to perform the conversion.
    * Prints the result or any errors encountered.

## Next Steps & Future Enhancements

This is a basic version. Here are some ideas for future improvements:

* **Fetch live exchange rates:**
    * Integrate with a free currency exchange rate API (e.g., [ExchangeRate-API](https://www.exchangerate-api.com/), or APIs from central banks if available and suitable for your use case).
    * This would involve making HTTP requests and parsing JSON responses.
* **More robust error handling:**
    * Handle invalid currency codes more gracefully.
    * Provide more specific error messages.
* **Complex conversion paths:**
    * If a direct rate (e.g., `AUD_CAD`) is not available, try to convert through a base currency (e.g., `AUD_USD` -> `USD_CAD`). This would require a more advanced rate storage structure or algorithm.
* **Configuration file:**
    * Allow loading rates or API keys from a configuration file instead of hardcoding.
* **Add more currencies:** Expand the list of supported currencies and their rates.
* **Unit tests:** Write tests for the `convertCurrency` function and input parsing logic.

---

Feel free to adjust any part of this to better fit your project's vision! This should give you a solid README to start with.