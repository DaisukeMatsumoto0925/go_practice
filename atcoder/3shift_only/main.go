package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		var in int
		fmt.Scan(&in)
		numbers[i] = in
	}

	b := newBlackBoard(numbers)
	b.countAll()

	fmt.Println(b.getCalculableCount())
}

type blackBoard struct {
	numbers         []int
	calculableCount int
}

func newBlackBoard(ns []int) *blackBoard {
	return &blackBoard{
		numbers: ns,
	}
}

func (b *blackBoard) isEvenAll() bool {
	for i := 0; i < len(b.numbers); i++ {
		if b.numbers[i]%2 == 0 {
			b.numbers[i] = b.numbers[i] / 2
		} else {
			return false
		}
	}
	return true
}

func (b *blackBoard) countUp() {
	b.calculableCount++
}

func (b *blackBoard) countAll() {
	for {
		if b.isEvenAll() {
			b.countUp()
		} else {
			break
		}
	}
}

func (b *blackBoard) getCalculableCount() int {
	return b.calculableCount
}
