package main

import (
	"fmt"
)

func main() {

	var n, a, b int
	fmt.Scan(&n)
	fmt.Scan(&a)
	fmt.Scan(&b)

	ary := []int{}

	for i := 0; i <= n; i++ {
		res := findSumOfDigits(n - i)
		if res >= a && res <= b {
			ary = append(ary, n-i)
		}
	}

	res := 0
	for _, v := range ary {
		res += v
	}

	fmt.Println(res)
}

func findSumOfDigits(n int) int {
	var sum int
	for {
		sum += n % 10
		n /= 10
		if n < 1 {
			break
		}
	}
	return sum
}
