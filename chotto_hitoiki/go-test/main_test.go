package main_test

import (
	"fmt"
	"testing"
)

// func setup() {
// 	fmt.Println("setup done")
// }

// func teardown() {
// 	fmt.Println("teardown done")
// }

// func TestMain(m *testing.M) {
// 	setup()

// 	m.Run()

// 	teardown()
// }

// func Test1(t *testing.T) {
// 	fmt.Println("test1 start")
// 	fmt.Println("test1 done")
// }

// func Test2(t *testing.T) {
// 	fmt.Println("test2 start")
// 	t.Fatal("test2 fatal!!")
// 	fmt.Println("test2 done")
// }

// func Test3(t *testing.T) {
// 	fmt.Println("test3 start")
// 	fmt.Println("test3 done")
// }

func TestTableDrivienParallel(t *testing.T) {
	t.Cleanup(func() {
		fmt.Println("cleanup done")
	})
	tests := []struct {
		name   string
		result bool
	}{
		{name: "subtest1", result: true},
		{name: "subtest2", result: true},
		{name: "subtest3", result: true},
	}

	for _, tt := range tests {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fmt.Printf("%s start\n", tc.name)

			if !tc.result {
				t.Fatalf("%s fatal!!\n", tc.name)
			}

			fmt.Printf("%s done\n", tc.name)
		})
	}
}
