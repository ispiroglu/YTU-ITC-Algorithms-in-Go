//Basic algorithm that add a number to another number.

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var sum int

	fmt.Println("Please enter the first number")
	fmt.Scan(&num1)

	fmt.Println("Please enter the second number")
	fmt.Scan(&num2)

	sum = num1 + num2

	fmt.Println("Sum is =", sum)
}
