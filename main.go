package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("*** Calculator ***")
	for {
		userChoice, err := userMenu()
		if err != nil {
			fmt.Println("Mistake! Choose the correct type of operation")
			continue
		}
		userData := getUserInput()
		userDataAll := stringNumbersInInt(userData)
		result := calculationData(userChoice, userDataAll)
		fmt.Printf("Result: %v\n", result)
		repeatCalculation := repeatCalculation()
		if !repeatCalculation {
			break
		}
	}
}

func userMenu() (string, error) {
	var userChoice string
	fmt.Print("Select the type of operation (AVG/SUM/MED): ")
	fmt.Scan(&userChoice)
	if userChoice == "AVG" || userChoice == "SUM" || userChoice == "MED" {
		return userChoice, nil
	}
	return "", errors.New("ERR_TYPE_CALCULATION")
}

func getUserInput() string {
	var userData string
	fmt.Print("Enter numbers for calculation separated by a comma: ")
	fmt.Scan(&userData)
	return userData
}

func stringNumbersInInt(userDataString string) []int {
	userDataStringSplit := strings.Split(userDataString, ",")
	var userDataInt []int
	for _, s := range userDataStringSplit {
		num, err := strconv.Atoi(s)
		if err == nil {
			userDataInt = append(userDataInt, num)
		}
	}
	return userDataInt
}

func calculationData(userChoice string, userDataInt []int) int {
	var resultCalculation int
	switch {
	case userChoice == "AVG":
		c := len(userDataInt)
		for i := 0; i < c; i++ {
			resultCalculation += (userDataInt[i])
		}
		resultCalculation = resultCalculation / c
	case userChoice == "SUM":
		for _, value := range userDataInt {
			resultCalculation = resultCalculation + value
		}
	case userChoice == "MED":
		l := len(userDataInt)
		if l == 0 {
			return 0
		} else if l%2 == 0 {
			resultCalculation = (userDataInt[l/2-1] + userDataInt[l/2]) / 2
		} else {
			resultCalculation = userDataInt[l/2]
		}
	}
	return resultCalculation
}

func repeatCalculation() bool {
	var userChoice string
	fmt.Print("Would you like to repeat the calculation? (Y/N): ")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	} else {
		return false
	}
}
