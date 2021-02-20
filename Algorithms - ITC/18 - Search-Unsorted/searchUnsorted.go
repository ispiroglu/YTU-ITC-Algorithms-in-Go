package main

import "fmt"

func main() {
	var n, x int
	var unsorted [50]int

	fmt.Println("Please enter the lenght of unsorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of unsorted array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&unsorted[i])
	}

	fmt.Println("Please enter the value that you're searcing for.")
	fmt.Scan(&x)

	i := 0
	for i < n && unsorted[i] != x {
		i++
	}

	if i < n {
		fmt.Println("Found at", i+1)
	} else {
		fmt.Println("Not found.")
	}
}
