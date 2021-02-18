//An algorihtm that checks the given number is Armstrong number.
//ERROR THIS ALGORITHM IS NOT WORKING RIGHT NOW. I'LL FIX IT SOON.
package main

import "fmt"

func main() {
	var num, reminder, digit, digitcount int
	digit = 0
	digitcount = 3

	fmt.Println("Please enter the number.")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a positive number.")
		return
	}
	var digits [3]int

	for i := 1; i <= digitcount; i++ {
		reminder = num
		dividecount := i - 1

		for j := 1; j <= dividecount; j++ {
			reminder = reminder / 10
			digits[digitcount-1] = digit
		}

	}

	sum := 0
	for i := 0; i < digitcount; i++ {
		digit = digits[i]
		sum = sum + (digit * digit * digit)
	}

	if sum == num {
		fmt.Print("That's an armstrong number.")
	} else {
		fmt.Print("X")
	}
	return

}
