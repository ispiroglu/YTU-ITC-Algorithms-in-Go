//An algorithm that finds the 3 digits Armstrong Numbers and count of them.
package main

import "fmt"

func main() {
	var i, count, digit1, digit2, digit3, cub1, cub2, cub3 int

	count = 0

	for i = 100; i <= 999; i++ {
		digit1 = i % 10
		digit2 = (i / 10) % 10
		digit3 = (i / 100)

		cub1 = digit1 * digit1 * digit1
		cub2 = digit2 * digit2 * digit2
		cub3 = digit3 * digit3 * digit3

		if i == cub1+cub2+cub3 {
			count++
			fmt.Println(i)
		}

	}
	fmt.Println("There are", count, "Armstrong numbers in 3 digits numbers")
}
