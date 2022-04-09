package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	checker := newIntHolder(a, b)
	caller := newCall()

	// 奇数チェック
	if checker.isOdd() {
		caller.printOdd()
		return
	}

	// 偶数チェック
	if checker.isEven() {
		caller.printEven()
		return
	}
}

type checker interface {
	isOdd() bool
	isEven() bool
}

type caller interface {
	printOdd()
	printEven()
}

type intHolder struct {
	a int
	b int
}

var _ checker = (*intHolder)(nil)

func newIntHolder(a, b int) *intHolder {
	return &intHolder{a, b}
}

func (i *intHolder) product() int {
	return i.a * i.b
}

func (i *intHolder) isEven() bool {
	res := i.product() % 2
	return res == 0
}

func (i intHolder) isOdd() bool {
	res := i.product() % 2
	return res != 0
}

var _ caller = (*call)(nil)

type call struct{}

func newCall() *call {
	return &call{}
}

func (c *call) printOdd() {
	fmt.Printf("Odd")
}

func (c *call) printEven() {
	fmt.Printf("Even")
}
