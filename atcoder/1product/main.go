package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	// インスタンス生成
	checker := newIntHolder(a, b)
	caller := newCall()

	// 奇数チェック
	if checker.isProductOdd() {
		caller.printOdd()
		return
	}

	// 偶数チェック
	if checker.isProductEven() {
		caller.printEven()
		return
	}

	panic(fmt.Sprintf("invalid args %d, %d,", a, b))
}

// checker check product of a and b is odd or even.
type checker interface {
	// isProductOdd return true when product of a and b is odd.
	isProductOdd() bool
	// isProductEven return true when product of a and b is even.
	isProductEven() bool
}

// caller print "Odd" or "Even"
type caller interface {
	// printOdd print "Odd"
	printOdd()
	// printEven print "Even"
	printEven()
}

// intHolder implement checker.
// create newIntHolder func.
type intHolder struct {
	a int
	b int
}

var _ checker = (*intHolder)(nil)

// newIntHolder constructor of intHolder address
func newIntHolder(a, b int) *intHolder {
	return &intHolder{a, b}
}

func (i *intHolder) product() int {
	return i.a * i.b
}

func (i *intHolder) isProductEven() bool {
	res := i.product() % 2
	return res == 0
}

func (i intHolder) isProductOdd() bool {
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
