//An algorithm that adds the column values of a matrix.

package main

import (
	"fmt"
)

func main() {
	var sum, row, column int
	var nums [20][20]int
	var sums [20]int

	fmt.Println("Please enter the amount of rows.")
	fmt.Scan(&row)

	fmt.Println("Please enter the amount of columns")
	fmt.Scan(&column)

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Println("Please enter the", i+1, ".row,", j+1, "column :")
			fmt.Scan(&nums[i][j])
		}
	}

	for j := 0; j < column; j++ {
		sum = 0
		for i := 0; i < row; i++ {
			sum += nums[i][j]
		}
		sums[j] = sum
	}

	for i := 0; i < column; i++ {
		fmt.Println("Sum of the", i+1, ". column is : ", sums[i])
	}

}
