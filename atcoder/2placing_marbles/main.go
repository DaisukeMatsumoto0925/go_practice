package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var in string
	fmt.Scan(&in)

	numbers := strings.Split(in, "")
	squareHolder, err := newSquareHolder(numbers)
	if err != nil {
		panic(err)
	}
	fmt.Println(squareHolder.count())
}

type Counter interface {
	count() int
}

var _ Counter = (*squareHolder)(nil)

type squareHolder struct {
	first  int
	second int
	third  int
}

func newSquareHolder(numbers []string) (*squareHolder, error) {
	if len(numbers) > 3 {
		return nil, errors.New("invalid digit. must be 3digit")
	}
	first, err := strconv.Atoi(numbers[0])
	if err != nil {
		return nil, err
	}
	second, err := strconv.Atoi(numbers[1])
	if err != nil {
		return nil, err
	}
	third, err := strconv.Atoi(numbers[2])
	if err != nil {
		return nil, err
	}
	return &squareHolder{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (s *squareHolder) count() int {
	var counts int
	if s.first == 1 {
		counts++
	}
	if s.second == 1 {
		counts++
	}
	if s.third == 1 {
		counts++
	}
	return counts
}
