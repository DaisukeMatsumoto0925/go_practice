package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	t := make([]int, N+1)
	x := make([]int, N+1)
	y := make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}

	for i := 0; i < N; i++ {
		dt := abs(t[i] - t[i+1])
		dx := abs(x[i] - x[i+1])
		dy := abs(y[i] - y[i+1])
		if !check(dt, dx+dy) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func check(dt, dist int) bool {
	return dist <= dt && dist%2 == dt%2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
