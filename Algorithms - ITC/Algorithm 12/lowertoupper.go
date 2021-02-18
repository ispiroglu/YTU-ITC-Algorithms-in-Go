//An algorithm that converts lower to upper.
//In this project i have decided to not use any other library other than "fmt" but i couldn't figure it how to change single characters in go.
//So i'll use strings library

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string

	fmt.Println("Please enter the word")
	fmt.Scan(&str1)

	res1 := strings.ToUpper(str1)
	fmt.Println(res1)

}
