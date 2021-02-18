//An algorithm that finding the elevation of lineer graphic that already given first and last point coordinations

package main

import "fmt"

func main() {
	var x1, x2, y1, y2 float64 = 0, 0, 0, 0

	var elevation float64

	fmt.Println("Please enter the coordinates like in order to x1,x2,y1,y2")
	fmt.Scan(&x1)
	fmt.Scan(&x2)
	fmt.Scan(&y1)
	fmt.Scan(&y2)

	if x1 == x2 && y1 == y2 {
		fmt.Println("You must enter two different coordinates")
		return
	}

	if x1 == x2 {
		fmt.Println("Elevation cannot be calculated properly.")
		return
	}

	elevation = (y2 - y1) / (x2 - x1)

	if elevation != 0 {
		fmt.Println("Elevation is ", elevation)
	}

	return
}
