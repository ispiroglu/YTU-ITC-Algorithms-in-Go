//An algorithm that turns the whole string to a number.

package main

import "fmt"

func main() {
	var i, number, tempNum, count int
	var word string

	fmt.Scan(&word)
	i = 0

	for range word {
		count++
	}

	fmt.Println(count)

	number = 0

	for i = count - 1; i >= 0; i-- {
		tempNum = int(word[i])
		number += tempNum
	}

	fmt.Println("The number is", number)

}
