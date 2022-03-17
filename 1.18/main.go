package main

import "fmt"

// Type parameter
func mapFunc[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a), cap(a))
	for i, e := range a {
		n[i] = f(e)
	}

	return n
}

func findFunc[T comparable](a []T, v T) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

// Type Sets
type Addable interface {
	int | int8 | int16 | int32 | int64 | ~uint64 | float64
	//  bool
}

func add[T Addable](a, b T) T {
	return a + b
}

// Constraints package
// 無くなった

// func main() {
// 	ai := []int{1, 2, 3, 4, 5, 6}
// 	as := mapFunc(ai, func(v int) string {
// 		return "<" + fmt.Sprint(v) + ">"
// 	})
// 	fmt.Println(as)
// }
// --------------------------------------------------

type Foo[T Addable] struct{ v T }

func (foo *Foo[T]) Add(v T) {
	foo.v += v
}

type Option[T Addable] func(*Foo[T])

func WithValue[T Addable](v T) Option[T] {
	return func(f *Foo[T]) { f.v = v }
}

func NewFoo[T Addable](options ...Option[T]) *Foo[T] {
	foo := new(Foo[T])
	for _, opt := range options {
		opt(foo)
	}
	return foo
}

func main() {
	foo := NewFoo(WithValue(3.14))
	foo.Add(4)

	fmt.Print(foo.v)
}
