//We are using a sorted array because when we find the value that bigger then our searching value, we are going to stop.

package main

import "fmt"

func main() {
	var sorted [50]int
	var n, x int

	fmt.Println("Please enter the lenght of sorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of sorted array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&sorted[i])
	}

	fmt.Println("Please enter the value that you're searcing for.")
	fmt.Scan(&x)

	i := 0
	for i < n && sorted[i] < x {
		i++
	}
	if i < n && sorted[i] == x {
		fmt.Println("Found at", i+1)
	} else {
		fmt.Println("Not found. Number of tries =", i+1)
	}

}
