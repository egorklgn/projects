package main

import (
	"errors"
	"fmt"
)

const USDtoEUR = 0.89
const USDtoRUB = 79.6
const EURtoRUB = USDtoRUB / USDtoEUR

func main() {
	fmt.Println("___ Currency converter app ___")
	for {
		originalCurrency, errOriginal := getUserInputOriginalCurrency()
		if errOriginal != nil {
			fmt.Print("Error! Currency must be either USD or EUR or RUB")
			break
		}
		amountForConversion, errAmount := getUserInputAmountForConversion()
		if errAmount != nil {
			fmt.Print("Error! The amount cannot be 0 or less than 0")
			break
		}
		targetCurrency, errTarget := getUserInputTargetCurrency(originalCurrency)
		if errTarget != nil {
			fmt.Print("Error! Please select a currency from those offered")
			break
		}
		result := calculateCurrency(originalCurrency, amountForConversion, targetCurrency)
		fmt.Printf("Total: %.2f\n", result)
		userChoice := repeatCalculation()
		if !userChoice {
			break
		}
	}

}

func getUserInputOriginalCurrency() (string, error) {
	var originalCurrency string
	fmt.Print("Enter the original currency (USD/EUR/RUB): ")
	fmt.Scan(&originalCurrency)
	if originalCurrency == "USD" || originalCurrency == "EUR" || originalCurrency == "RUB" {
		return originalCurrency, nil
	} else {
		return "", errors.New("INVALID")
	}
}

func getUserInputAmountForConversion() (float64, error) {
	var amountForConversion float64
	fmt.Print("Enter the amount to convert: ")
	fmt.Scan(&amountForConversion)
	if amountForConversion <= 0.0 {
		return 0.0, errors.New("INVALID")
	} else {
		return amountForConversion, nil
	}
}

func getUserInputTargetCurrency(originalCurrency string) (string, error) {
	var targetСurrency string
	switch originalCurrency {
	case "USD":
		fmt.Print("Enter the target currency (EUR/RUB): ")
	case "EUR":
		fmt.Print("Enter the target currency (USD/RUB): ")
	case "RUB":
		fmt.Print("Enter the target currency (USD/EUR): ")
	}
	fmt.Scan(&targetСurrency)
	if targetСurrency == "USD" || targetСurrency == "EUR" || targetСurrency == "RUB" {
		return targetСurrency, nil
	} else {
		return "", errors.New("INVALID")
	}
}

func calculateCurrency(originalCurrency string, amountForConversion float64, targetCurrency string) float64 {
	var result float64
	switch {
	case originalCurrency == "USD" && targetCurrency == "EUR":
		result = amountForConversion * USDtoEUR
	case originalCurrency == "USD" && targetCurrency == "RUB":
		result = amountForConversion * USDtoRUB
	case originalCurrency == "EUR" && targetCurrency == "USD":
		result = amountForConversion / USDtoEUR
	case originalCurrency == "EUR" && targetCurrency == "RUB":
		result = amountForConversion * EURtoRUB
	case originalCurrency == "RUB" && targetCurrency == "USD":
		result = amountForConversion / 79.6
	case originalCurrency == "RUB" && targetCurrency == "EUR":
		result = amountForConversion / EURtoRUB
	}
	return result
}

func repeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to repeat the conversion? (Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
