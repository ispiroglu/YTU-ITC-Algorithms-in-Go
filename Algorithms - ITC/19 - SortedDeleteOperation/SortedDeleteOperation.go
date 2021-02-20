// An algorithm that deletes a value from sorted array.
package main

import "fmt"

func main() {
	var n, x int
	var arr [50]int

	fmt.Println("Please enter the lenght of sorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of sorted array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Please enter the value you want to remove")
	fmt.Scan(&x)

	i := 0
	for i < n && arr[i] < x {
		i++
	}

	if arr[i] == x {
		for j := i; j < n-1; j++ {
			arr[j] = arr[j+1]
		}
		n--
	} else {
		fmt.Println("The value is not in this array.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])

		if i != n-1 {
			fmt.Print(",")
		}
	}

}
