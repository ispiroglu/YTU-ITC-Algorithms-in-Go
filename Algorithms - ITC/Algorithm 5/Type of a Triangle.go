//An algorithm that finds the type of the triangle by its edges

package main

import "fmt"

func main() {
	var e1, e2, e3, biggest, d1, d2 float64

	fmt.Println("Please enter the edges of triangle in order to Egde 1, Edge 2, Edge 3.")
	fmt.Scan(&e1)
	fmt.Scan(&e2)
	fmt.Scan(&e3)

	if e1 > e2 {
		biggest = e1
		d1 = e2
		d2 = e3
	} else {
		biggest = e2
		d1 = e1
		d2 = e3
	}

	if e3 > e2 && e3 > e1 {
		biggest = e3
		d1 = e1
		d2 = e2
	}

	fmt.Println("Biggest is : ", biggest)

	if e1 == e2 {
		if e1 == e3 {
			fmt.Println("Triangle is a equilateral triangle.")
		} else if (biggest*biggest)-((d1*d1)+(d2*d2)) <= 0.00001 {
			fmt.Println("Triangle is a right triangle")
		} else {
			fmt.Println("Triangle is a isosceles triangle")
		}
	} else {
		if (biggest*biggest)-((d1*d1)+(d2*d2)) <= 0.00001 {
			fmt.Println("Triangle is a right triangle")
		} else if e3 == e2 {
			fmt.Println("Triangle is a isosceles triangle")
		} else {
			fmt.Println("Triangle is a scalane tirangle")
		}
	}
}
