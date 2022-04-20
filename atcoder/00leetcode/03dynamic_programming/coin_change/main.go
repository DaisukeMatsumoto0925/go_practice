package main

// func main() {
// 	// fmt.Println(coinChange([]int{1, 2, 5}, 11))
// 	// fmt.Println(coinChange([]int{1, 2147483647}, 2))
// 	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
// }

// func coinChange(coins []int, amount int) int {
// 	if amount == 0 {
// 		return 0
// 	}

// 	sum := 0
// 	for _, c := range coins {
// 		sum += c
// 	}

// 	var res int
// 	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
// 	for _, c := range coins {
// 		for c <= amount {
// 			amount -= c
// 			res++
// 		}
// 	}

// 	if amount != 0 {
// 		return -1
// 	}

// 	return res
// }
