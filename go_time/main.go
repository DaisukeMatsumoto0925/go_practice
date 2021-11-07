package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// tz, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Println(time.Date(2019, 8, 28, 23, 0, 0, 0, time.UTC))

	const format1 = "2006-01-02 15:04:05" // YYYY-MM-DD hh:mm:ss表記
	const format2 = "2006年01月02日"         // 漢字表記／年月日のみ
	const format3 = "15時04分"              // 漢字表記／時分のみ
	const format4 = "15:04"               // 漢字表記／時分のみ

	t1, _ := time.Parse(format1, "2019-08-28 22:30:00")
	t2, _ := time.Parse(format2, "2019年08月28日")
	t3, _ := time.Parse(format3, "22時30分")
	t4, _ := time.Parse(format4, "22:30")

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
	fmt.Println(t4)

}
