package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 3番目
func main() {
	sc.Split(bufio.ScanWords)
	cap := 6
	ary := make([]int, cap)
	for i := 0; i < cap; i++ {
		ary[i] = nextInt()
	}
	sort.SliceStable(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})
	fmt.Println(ary[3])
}

//増減管理
// func main() {
// 	var n string
// 	if sc.Scan() {
// 		n = sc.Text()
// 	}
// 	nInt, _ := strconv.Atoi(n)
// 	dailySales := make([]int, nInt)
// 	for i := 0; i < cap(dailySales); i++ {
// 		var sale string
// 		if sc.Scan() {
// 			sale = sc.Text()
// 		}
// 		saleInt, _ := strconv.Atoi(sale)

// 		dailySales[i] = saleInt
// 	}

// 	for i := 0; i < cap(dailySales)-1; i++ {
// 		var diff int
// 		diff = dailySales[i+1] - dailySales[i]
// 		if diff > 0 {
// 			fmt.Printf("up %v\n", math.Abs(float64(diff)))
// 		} else if diff < 0 {
// 			fmt.Printf("down %v\n", math.Abs(float64(diff)))
// 		} else {
// 			fmt.Println("stay")
// 		}
// 	}
// }

// 2倍チェック
// func main() {
// 	var s string
// 	if sc.Scan() {
// 		s = sc.Text()
// 	}

// 	n, err := strconv.Atoi(s)
// 	if err != nil {
// 		fmt.Println("error")
// 	} else {
// 		fmt.Println(n * 2)
// 	}
// }

// func main() {
// 	var a int
// 	var b int

// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	fmt.Println(a * b)
// }

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Scan()
// 	var n, _ = strconv.Atoi(sc.Text())
// 	for i := 0; i < n; i++ {
// 		sc.Scan()
// 		var s = strings.Split(sc.Text(), " ")
// 		fmt.Println("hello = " + s[0] + " , world = " + s[1])
// 	}
// }
