package main

import (
	"fmt"
	"strconv"
)

type Car interface {
	run(int) string
	stop()
}

type MyCar struct {
	name  string
	speed int
}

// run
func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

// stop
func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func main() {
	myCar := &MyCar{
		name:  "マイカー",
		speed: 0,
	}

	// MyCarの構造体を持ち、Carインターフェースを持つ変数の定義
	var objCar Car = myCar
	fmt.Println(objCar)
	fmt.Println(objCar.run(50))
	objCar.stop()
}
