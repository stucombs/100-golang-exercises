package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but
are not a multiple of 5, between 2000 and 3200 (both included). The numbers
obtained should be printed in a comma-separated sequence on a single line.

Hint: consider using strconv and strings.Join.

Tip: run `go test ./...` from this folder.
*/

// Ex001 should return all integers in [low, high] that are divisible by 7
// but not by 5, joined by commas (e.g. "112,119,126,...").

func main() {
	res := ex001(2000, 3200)

	fmt.Println(res)
}

func ex001(low, high int) string {
	var numbers []string
	i := low
	for i <= high {
		if i % 7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
		i++
	}

	return strings.Join(numbers, ",")
}
