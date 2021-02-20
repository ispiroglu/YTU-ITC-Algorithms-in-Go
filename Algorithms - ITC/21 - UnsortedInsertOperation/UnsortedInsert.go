package main

import "fmt"

func main() {
	var x, n, index int
	var arr [50]int

	fmt.Println("Please enter the lenght of sorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of sorted array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Please enter the value you want to insert")
	fmt.Scan(&x)
	fmt.Println("Please enther the index you want to insert to array in")
	fmt.Scan(&index)
	index--

	i := 0
	for i < n && arr[i] < x {
		i++
	}

	for i = n - 1; i >= index; i-- {
		arr[i+1] = arr[i]
	}

	arr[index] = x
	n++

	for i := 0; i < n; i++ {
		fmt.Print(arr[i])

		if i != n-1 {
			fmt.Print(",")
		}
	}
}
