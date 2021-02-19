//An algorithm that finds are there any repeating numbers in the list.

package main

import "fmt"

func main() {
	var i, j, n, found int
	var nums [50]int

	fmt.Println("Enter the length of the array")
	fmt.Scan(&n)

	fmt.Println("Now please enter the numbers.")
	for i = 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	i = 0
	found = 0

	for i < n && found != 1 {
		j = i + 1
		for j < n && found != 1 {
			if nums[i] == nums[j] {
				found = 1
			}
			j++
		}
		i++
	}

	if found == 0 {
		fmt.Println("There is no repeating numbers")
	} else {
		fmt.Println("There is repeatition")
	}

}
