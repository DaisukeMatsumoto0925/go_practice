package main

import (
	"fmt"
	"sort"
)

func main() {
	var count int
	fmt.Scan(&count)

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		var n int
		fmt.Scan(&n)
		numbers[i] = n
	}

	// var first, second int
	sort.Ints(numbers)

	var first, second int
	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			first += numbers[i]
		} else {
			second += numbers[i]
		}
	}

	fmt.Print(first - second)
	fmt.Print(first, second)
}
