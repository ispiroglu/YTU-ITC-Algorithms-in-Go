//Palindrome checker

package main

import "fmt"

func main() {
	var n, cursor1, cursor2 int
	var word string

	fmt.Println("Please enter your word")
	fmt.Scan(&word)

	for range word {
		n++
	}

	cursor1 = 0
	cursor2 = n - 1

	for cursor1 < cursor2 && word[cursor1] == word[cursor2] {
		cursor1++
		cursor2--
	}

	if cursor1 >= cursor2 {
		fmt.Println("That's a Palindrome word.")
	} else {
		fmt.Println("That's not a Palindrome word")
	}
}
