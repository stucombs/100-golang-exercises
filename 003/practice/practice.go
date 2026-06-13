package main

import (
	"fmt"
	"log"
)

/*
Exercise 003:

With a given integral number n, write a program to generate a map that contains
(i, i*i) such that i is an integral number between 1 and n (both included).

Example: Ex003(8) -> map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

Tip: run `go test ./...` from this folder.
*/

func main() {
	var input int
	fmt.Print("Input a number: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Error occured : ", err)
	}
	res := Ex003(input)

	fmt.Printf("Mapping for %d is %v\n", input, res)
}

// Ex003 should return a map where keys are 1..n and values are key*key.
func Ex003(max int) map[int]int {
	mapping := make(map[int]int)

	if max <= 0 {
		log.Fatal("Number below 1 not allowed")
	}

	for n := 1; n <= max; n++ {
		mapping[n] = n * n
	}

	return mapping
}
