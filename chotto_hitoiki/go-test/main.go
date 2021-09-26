package main

import "fmt"

func main() {
	array := []int{1, 2, 3}
	fmt.Println(BisectLeft(array, 1))
}

func BisectLeft(array []int, key int) int {
	var left, right, middle int
	left = 0
	right = len(array)
	for left < right {
		middle = ((right - left) / 2) + left
		if key < array[middle] {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}
