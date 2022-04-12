package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	nary := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nary[i])
	}

	sort.Ints(nary)
	var res int
	for i := 0; i < n; i++ {
		if i == 0 {
			res++
			continue
		}

		if nary[i] != nary[i-1] {
			res++
		}
	}

	fmt.Println(res)
}
