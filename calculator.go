package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	// User input for first number
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Fatal("Invalid input")
		os.Exit(1)
	}

	// User input for operator
	fmt.Print("Enter an operator (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatal("Invalid input")
		os.Exit(1)
	}

	// User input for second number
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatal("Invalid input")
		os.Exit(1)
	}

	var result float64

	// Perform the operation based on the operator
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			log.Fatal("Error: Division by zero")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		log.Fatal("Invalid operator. Use +, -, *, or /.")
		os.Exit(1)
	}

	// Output the result
	fmt.Printf("The result of %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
