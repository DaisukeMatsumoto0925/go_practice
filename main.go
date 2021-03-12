package main

import (
	"bufio"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

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
