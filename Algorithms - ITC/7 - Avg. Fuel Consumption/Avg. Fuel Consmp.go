//An algorithm that calculates the average fuel consumption. In every second, user will enter the fuel consumption.
//Value "-1" will mean car has stopped.

package main

import "fmt"

func main() {
	var time, consumption, total, average float64 = 0, 0, 0, 0

	fmt.Println("Please enter the consumption of car in second")
	fmt.Scan(&consumption)

	for consumption > 0 {
		total = total + consumption
		time++
		average = total / time
		fmt.Println("Consumption is : ", consumption)
		fmt.Println("Please enter the consumption of car in second")
		fmt.Scan(&consumption)
	}

	fmt.Println("For", time, "seconds average is", average)

}
