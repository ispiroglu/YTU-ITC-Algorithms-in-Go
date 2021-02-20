package main

import "fmt"

func main() {
	var n, m int
	var arr1, arr2 [100]int
	var sorted [200]int

	fmt.Println("Please enter the lenght of first sorted array.")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of first array.")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Println("Please enter the lenght of second sorted array.")
	fmt.Scan(&m)

	fmt.Println("Please enter the values of second array.")
	for i := 0; i < m; i++ {
		fmt.Scan(&arr2[i])
	}

	i := 0
	j := 0

	for j < m && i < n {
		if arr1[i] < arr2[j] {
			sorted[i+j] = arr1[i]
			i++
		} else {
			sorted[i+j] = arr2[j]
			j++
		}
	}

	if j < m {
		for j < m {
			sorted[j+i] = arr2[j]
			j++
		}
	} else {
		for i < n {
			sorted[j+i] = arr1[i]
			i++
		}
	}

	for i := 0; i < m+n; i++ {
		fmt.Print(sorted[i])

		if i != m+n-1 {
			fmt.Print(",")
		}
	}

}
