package main

import "fmt"

func main() {
	CLIInterface()
}

func CLIInterface() {
	var userChoice int
	fmt.Println("Welcome to the calculator powered by GO")
	fmt.Println("1. Add numbers \n2. Substract number \n3. Multiply numbers \n4. Divide numbers \n5. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&userChoice)
	if userChoice == 5 {
		fmt.Println("Thank you for using our services.")
	} else {
		fmt.Println("Enter the two numbers")
		var num1, num2 = takeUserInput()
		var result int = calculation(num1, num2, userChoice)
		if result == 0 {
			fmt.Println("Thank you for using our services.")
		} else {
			fmt.Printf("The result is %v", result)
		}
	}
}

func takeUserInput() (int, int) {
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	return num1,num2
}

func calculation(num1, num2, userChoice int) int {
	switch userChoice{
	case 1:
		return num1+num2
	case 2:
		return num1-num2
	case 3: 
		return num1*num2
	case 4:
		return num1/num2
	case 5:
		break
	default:
		fmt.Println("Enter a valid choice")
	}
	return 0
}

