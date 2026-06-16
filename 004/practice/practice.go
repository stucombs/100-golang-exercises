package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Exercise 004:

Write a program which accepts a sequence of comma-separated numbers and
generates a slice out of them.

Example:
  Ex004("12, 23, 34, 45") -> []int{12, 23, 34, 45}
*/

func main() {
	var input string = "1, 10, 25, 45"
	res := Ex004(input)

	fmt.Print(res)
}

// Ex004 should parse a comma-separated list of integers (whitespace is allowed
// around the numbers) and return them as []int.
func Ex004(input string) []int {
	strArr := strings.Split(input, ",")
	result := []int{}

	for _, n := range strArr {
		n = strings.Trim(n, " ")
		conv, err := strconv.Atoi(n)

		if err != nil {
			log.Fatal("Failed to convert string to int")
		}

		result = append(result, conv)
	}

	return result
}
