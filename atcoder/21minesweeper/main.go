package main

import (
	"fmt"
)

func main() {
	hoge := "hoge"
	fmt.Println(hoge[1])
	// var h, w int
	// fmt.Scan(&h, &w)
	// s := make([]string, h)
	// for i := range s {
	// 	fmt.Scan(&s[i])
	// }
	h := 3
	w := 5
	s := []string{".....", ".#.#.", "....."}

	a := []int{-1, 0, 1}
	// 行数
	for i := range s {
		// 列数
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				fmt.Print("#")
				continue
			}
			r := 0
			for _, p := range a {
				for _, q := range a {
					b, c := i+p, j+q
					if (p != 0 || q != 0) && 0 <= b && b < h && 0 <= c && c < w {
						if s[b][c] == '#' {
							r++
						}
					}
				}
			}
			fmt.Print(r)
		}
		fmt.Println()
	}
}
