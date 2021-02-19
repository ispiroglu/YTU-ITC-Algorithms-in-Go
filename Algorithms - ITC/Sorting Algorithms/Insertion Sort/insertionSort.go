//Basic Insertion sort.

package main

import "fmt"

func main() {
	var arr [100]int
	var n, tmp int

	fmt.Println("Please enter the lenght of array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 1; i < n; i++ {
		j := i - 1
		tmp = arr[i]

		for j >= 0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp

	}

	fmt.Println("Sorted.")

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])

		if i != n-1 {
			fmt.Print(",")
		}
	}

}
