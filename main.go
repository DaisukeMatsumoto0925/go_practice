package main

// func main() {
// 	var n, q int
// 	fmt.Scan(&n, &q)
// 	followList := make([][]int, n)
// 	for i := range followList {
// 		followList[i] = make([]int, n)
// 	}

// 	for i := 0; i < q; i++ {
// 		var m, a, b int
// 		fmt.Scan(&m, &a)
// 		if m == 1 {
// 			fmt.Scan(&b)
// 			followList[a-1][b-1] = 1
// 		} else if m == 2 {
// 			for j := 0; j < n; j++ {
// 				if a-1 == j {
// 					continue
// 				}
// 				if followList[j][a-1] > 0 {
// 					followList[a-1][j] = 1
// 				}
// 			}
// 		} else if m == 3 {
// 			var addList []int
// 			for j := 0; j < n; j++ {
// 				if followList[a-1][j] > 0 {
// 					addList = append(addList, j)
// 				}
// 			}
// 			for _, j := range addList {
// 				for k := 0; k < n; k++ {
// 					if a-1 == k {
// 						continue
// 					}
// 					if followList[j][k] > 0 {
// 						followList[a-1][k] = 1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if followList[i][j] > 0 {
// 				fmt.Print("Y")
// 			} else {
// 				fmt.Printf("N")
// 			}
// 		}
// 		fmt.Print("\n")
// 	}
// }

// 重複チェック
// var sc = bufio.NewScanner(os.Stdin)

// func nextLine() int {
// 	sc.Scan()
// 	i, e := strconv.Atoi(sc.Text())
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// func main() {
// 	n := nextLine()
// 	ary := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		a := nextLine()
// 		ary[a] = ary[a] + 1
// 	}

// 	var (
// 		x int
// 		y int
// 	)
// 	for i, v := range ary {
// 		if i > 0 {
// 			if v > 1 {
// 				x = i
// 			}
// 			if v == 0 {
// 				y = i
// 			}
// 		}
// 	}

// 	if x > 0 {
// 		fmt.Printf("%d %d \n", x, y)
// 	} else {
// 		fmt.Println("Correct")
// 	}
// }

// var sc = bufio.NewScanner(os.Stdin)

// func nextInt() int {
// 	sc.Scan()
// 	i, e := strconv.Atoi(sc.Text())
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// // 3番目
// func main() {
// 	sc.Split(bufio.ScanWords)
// 	cap := 6
// 	ary := make([]int, cap)
// 	for i := 0; i < cap; i++ {
// 		ary[i] = nextInt()
// 	}
// 	sort.SliceStable(ary, func(i, j int) bool {
// 		return ary[i] < ary[j]
// 	})
// 	fmt.Println(ary[3])
// }

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
