//An algorithm that finds the intersection of 2 array.

package main

import "fmt"

func main() {
	var i, j, n, m, intersectionCount int
	var inSecSet [50]int
	var set2 [50]int
	var set1 [50]int

	fmt.Println("Please enter the lenght of first set")
	fmt.Scan(&n)

	fmt.Println("Please enter the values of first set")
	for i = 0; i < n; i++ {
		fmt.Scan(&set1[i])
	}

	fmt.Println("Please enter the lenght of second set")
	fmt.Scan(&m)

	fmt.Println("Please enter the values of second set")
	for i = 0; i < m; i++ {
		fmt.Scan(&set2[i])
	}

	for i = 0; i < n; i++ {
		j = 0

		for j < m && set1[i] != set2[j] {
			j++
		}
		if j < m {
			inSecSet[intersectionCount] = set1[i]
			intersectionCount++
		}
	}

	if intersectionCount != 0 {
		if intersectionCount == 1 {
			fmt.Print("The same number is : ", inSecSet[0])

		} else {
			fmt.Print("The same numbers are : ")
			for i = 0; i < intersectionCount; i++ {
				fmt.Print(inSecSet[i])
			}
		}
	} else {
		fmt.Print("There is no similarity between these two set.")
	}
}
