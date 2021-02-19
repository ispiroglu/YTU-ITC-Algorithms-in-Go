//Algorithm 2: Really basic a calculator.
package main

import "fmt"

func main() {
	var num1, num2, result int
	var operation string

	fmt.Println("Please enter your first operand.")
	fmt.Scan(&num1)

	fmt.Println("Please enter your second operand.")
	fmt.Scan(&num2)

	fmt.Println("Please enter your operation.")
	fmt.Scan(&operation)

	if operation == "+" {
		result = num1 + num2
	} else if operation == "-" {
		result = num1 - num2
	} else if operation == "*" {
		result = num1 * num2
	} else if operation == "/" {
		result = num1 / num2
	} else {
		fmt.Println("You have entered a wrong operation.")
		return
	}

	fmt.Println("Result = ", result)
	return

}
