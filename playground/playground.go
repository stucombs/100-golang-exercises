package main

import (
	"fmt"
)

func main() {
	number := 8

	var result int
	for number > 0 {
		// fmt.Println(number)
		result = result * number
		number--
	}

	fmt.Println(result)
	// var numbers []string
	// i := 2000
	// for i <= 3200 {
	// 	if i % 7 == 0 && i%5 != 0 {
	// 		numbers = append(numbers, strconv.Itoa(i))
}
