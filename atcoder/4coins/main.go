package main

import "fmt"

func main() {
	var fiveHCCounts, HCCounts, fifHCCounts, amount int
	fmt.Scan(&fiveHCCounts)
	fmt.Scan(&HCCounts)
	fmt.Scan(&fifHCCounts)
	fmt.Scan(&amount)

	var res int

	for i := 0; i <= fiveHCCounts; i++ {
		for j := 0; j <= HCCounts; j++ {
			for k := 0; k <= fifHCCounts; k++ {
				if i*500+j*100+k*50 == amount {
					res += 1
				}
			}
		}
	}

	fmt.Println(res)
}
