package main

import "fmt"

func main() {
	var operation string
	var firstValue int
	var secondValue int

	fmt.Println("Calculator GO 1.0")
	fmt.Println("*****************")
	fmt.Println("Enter the operation! (Add, Subtract, Multiply, Divide)")
	fmt.Scan(&operation)
	fmt.Println("Enter the first value")
	fmt.Scan(&firstValue)
	fmt.Println("Enter the second value")
	fmt.Scan(&secondValue)

	switch operation {
	case "Add":
		fmt.Println("Result:", firstValue+secondValue)
	case "Subtract":
		fmt.Println("Result:", firstValue-secondValue)
	case "Multiply":
		fmt.Println("Result:", firstValue*secondValue)
	case "Divide":
		if secondValue != 0 {
			fmt.Println("Result:", firstValue/secondValue)
		} else {
			fmt.Println("Error: Cannot divide by zero")
		}
	default:
		fmt.Println("Invalid Operation")
	}
}
