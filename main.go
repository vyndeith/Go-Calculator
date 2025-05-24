package main

import (
	"fmt"
	"strconv"
)

func main() {
	firstNumber := inputNumber("First number: ")
	secondNumber := inputNumber("Second number: ")
	fmt.Print("Select operation (+, -, *, /): ")
	var operation string
	fmt.Scanln(&operation)
	fmt.Printf("You wrote: %2f %s %2f\n", firstNumber, operation, secondNumber)
	if operation == "+" {
		fmt.Printf("Result: %2f\n", firstNumber+secondNumber)
	} else if operation == "-" {
		fmt.Printf("Result: %2f\n", firstNumber-secondNumber)
	} else if operation == "*" {
		fmt.Printf("Result: %2f\n", firstNumber*secondNumber)
	} else if operation == "/" {
		if secondNumber == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		fmt.Printf("Result: %f\n", firstNumber/secondNumber)
	} else {
		fmt.Println("Unknown operation")
	}
}

func inputNumber(prompt string) float64 {
	var number float64
	for {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)

		var err error
		number, err = strconv.ParseFloat(input, 64)
		if err == nil {
			return number
		}
		fmt.Println("Invalid input. Please enter a number.")
	}
}
