package main

import (
	"fmt"

	"github.com/DaisukeMatsumoto0925/go-test/bisect"
)

func main() {
	array := []int{1, 2, 3}
	fmt.Println(bisect.BisectLeft(array, 1))
}
