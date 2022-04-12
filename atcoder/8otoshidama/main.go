package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	for a := 0; a <= n; a++ {
		for b := 0; b+a <= n; b++ {
			c := n - a - b
			total := 10000*a + 5000*b + 1000*c
			if total == y {
				fmt.Print(a, b, c)
				return
			}
		}
	}

	fmt.Print(-1, -1, -1)
}
