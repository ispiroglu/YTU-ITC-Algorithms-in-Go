package main

import "fmt"

func main() {
	var tmp, min, n int
	var arr [50]int

	fmt.Println("Please enter the lenght of array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of array")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp

	}

	fmt.Println("Sorted.")

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
		if i != n-1 {
			fmt.Print(",")
		}
	}

}
