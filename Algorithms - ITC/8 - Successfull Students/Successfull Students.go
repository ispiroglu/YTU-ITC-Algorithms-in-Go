//An algorithm that finds the who passed the finals by looking the average of class.

package main

import "fmt"

func main() {
	var average float64
	var i, n, sum int
	var notes [100]int

	fmt.Println("How many students in class ?")
	fmt.Scan(&n)

	fmt.Println("Enter the notes please.")
	for i = 0; i < n; i++ {
		fmt.Scan(&notes[i])
		sum = sum + notes[i]
	}

	average = float64(sum / n)

	for i = 0; i < n; i++ {
		if float64(notes[i]) >= average {
			fmt.Println(i+1, "listed person have passed the class")
		}
	}

}
