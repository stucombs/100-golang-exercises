package main

import (
	"fmt"
	"log"
)

/*
Exercise 002:

Write a program which can compute the factorial of a given number.

  - Ex002(8) -> 40320, nil
  - Ex002(0) -> 1,     nil   (by definition)
  - Ex002(-3) -> 0,    error (negative input is not allowed)

Tip: run `go test ./...` from this folder.
*/

func main() {
	var input int
	fmt.Print("Input a number: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	res, err := Ex002(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Factorial of %d = %d\n", input, res)
}

// Ex002 should return n! as a uint64 (or an error for negative input).
func Ex002(input int) (uint64, error) {
	var factorial uint64 = 1

	if input < 0 {
		return 0, fmt.Errorf("Negative number input")
	}

	if input == 0 {
		return 1, nil
	}

	for n := 1; n <= input; n++ {
		factorial *= uint64(n)
	} 

	return factorial, nil
}
