//An algorithm that founds the coldest and the hottest day of a year.
//Temperature values will be entered by user.
//Tested with lower numbers.

package main

import "fmt"

func main() {
	var i, coldest, hottest int
	var temps [365]int

	fmt.Println("Please enter the temperatures of every day")

	for i = 0; i < 365; i++ {
		fmt.Scan(&temps[i])
	}

	hottest = temps[0]
	coldest = temps[1]

	for i = 1; i < 365; i++ {
		if hottest < temps[i] {
			hottest = temps[i]
		} else if coldest > temps[i] {
			coldest = temps[i]
		}
	}

	fmt.Println("The hottest is :", hottest, "The coldest is :", coldest)
}
