package grasshopper

import "fmt"

type HighJumpRunner interface {
	HighJump() string
	Run() string
}

type grasshopper struct{}

func (g *grasshopper) HighJump() string{
	return "High Jump!"
}

type ShinJinrui2 struct {
	*grasshopper
}

func InitGrasshopper() {
	aGrassHopper := &grasshopper{}
	aShinJinrui2 := &ShinJinrui2{
		grasshopper: aGrassHopper,
	}
	if _, ok := interface{}(aShinJinrui2).(HighJumpRunner); ok {
        fmt.Println("ShinJinrui2はHighJumpRunnerインターフェースを実装しています。")
    }
	fmt.Println(aShinJinrui2.HighJump())
}
