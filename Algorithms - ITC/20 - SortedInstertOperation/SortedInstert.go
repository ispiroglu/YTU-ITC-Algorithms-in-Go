package main

import "fmt"

func main() {
	var n, x, j int
	var sorted [50]int

	fmt.Println("Please enter the lenght of sorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of sorted array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&sorted[i])
	}

	fmt.Println("Please enter the value you want to insert")
	fmt.Scan(&x)

	i := 0
	for i < n && sorted[i] < x {
		i++
	}

	for j = n - 1; j >= i; j-- {
		sorted[j+1] = sorted[j]
	}
	sorted[i] = x
	n++

	for i := 0; i < n; i++ {
		fmt.Print(sorted[i])

		if i != n-1 {
			fmt.Print(",")
		}
	}
}
