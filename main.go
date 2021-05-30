package main

import (
	"github.com/DaisukeMatsumoto0925/go_embedding/dolphin"
)

type Flyer interface {
	Fly() string
}

type Runner interface {
	Run() string
}

type FlyingRunner interface {
	Flyer
	Runner
}

type ToriJin struct {
	FlyingRunner
}

type ShinJinrui struct {
	*ToriJin
}

type RealToriJin struct{}

func (r RealToriJin) Fly() string {return "Fly!"}

func (r RealToriJin) Run() string {return "Run!"}

func main() {
	// aRealToriJin := &RealToriJin{}
	// aTorijin := &ToriJin{
	// 	FlyingRunner: aRealToriJin,
	// }
	// aShinJinrui := &ShinJinrui{
	// 	ToriJin: aTorijin,
	// }
	// fmt.Println(aShinJinrui.Fly())
	// fmt.Println(aShinJinrui.Run())
	// grasshopper.InitGrasshopper()
	// status.InitStatus()
	dolphin.InitDolphin()
}
