package main

import "fmt"

func main() {
	const USDtoEUR = 0.89
	const USDtoRUB = 79.6
	const EURtoRUB = USDtoRUB / USDtoEUR
}

func getUserInput() (string, float64, string) {
	var originalCurrency string
	var amountForConversion float64
	var targetСurrency string
	fmt.Print("Enter the original currency: ")
	fmt.Print("Enter the amount to convert: ")
	fmt.Print("Enter the target currency: ")
	fmt.Scan(&originalCurrency)
	fmt.Scan(&amountForConversion)
	fmt.Scan(&targetСurrency)
	return originalCurrency, amountForConversion, targetСurrency
}

func calculateCurrency(originalCurrency string, amountForConversion float64, targetCurrency string) float64 {

}
