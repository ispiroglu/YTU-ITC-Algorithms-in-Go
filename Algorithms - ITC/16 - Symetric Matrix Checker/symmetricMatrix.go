//An algorithm that checks the martix is symetric or not.

package main

import "fmt"

func main() {
	var n int
	var matrix [50][50]int

	fmt.Println("For checking symetry, Matrix must be a square matrix.")

	fmt.Println("Please enter the dimenson of matrix.")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println("Please enter the", i+1, ". row,", j+1, "column value.")
			fmt.Scan(&matrix[i][j])
		}
	}

	i := 0
	symetric := 1
	for i < (n-1) && symetric == 1 {
		j := i + 1
		for j < n && symetric == 1 {
			if matrix[i][j] != matrix[j][i] {
				symetric = 0
			}
			j++
		}
		i++
	}

	if symetric == 1 {
		fmt.Println("Symetric.")
	} else {
		fmt.Println("Not symetric.")
	}

}
