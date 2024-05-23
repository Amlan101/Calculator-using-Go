package main

import (
	"fmt"
	"os"
)

func main() {
	CLIInterface()
}

func CLIInterface() {
	for {
		userChoice := getUserChoice()
		if userChoice == 5 {
			fmt.Println("Thank you for using our services.")
			break
		}
		num1, num2 := takeUserInput()
		result, err := calculation(num1, num2, userChoice)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("The result is: %.2f\n", result)
		}
	}
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Welcome to the calculator powered by GO")
	fmt.Println("1. Add numbers")
	fmt.Println("2. Subtract numbers")
	fmt.Println("3. Multiply numbers")
	fmt.Println("4. Divide numbers")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&userChoice)
	return userChoice
}

func takeUserInput() (float64, float64) {
	var num1, num2 float64
	fmt.Print("Enter the two numbers separated by space: ")
	_, err := fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println("Invalid input. Please enter two numbers.")
		os.Exit(1)
	}
	return num1, num2
}

func calculation(num1, num2 float64, userChoice int) (float64, error) {
	switch userChoice {
	case 1:
		return num1 + num2, nil
	case 2:
		return num1 - num2, nil
	case 3:
		return num1 * num2, nil
	case 4:
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid choice")
	}
}

