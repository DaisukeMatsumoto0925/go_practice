package dolphin

import "fmt"

type dolphin struct{}

func (g *dolphin) Dive() string{
	return "Dolpin Dive!"
}

type ShinJinrui4 struct {
	*dolphin
}

func (sj *ShinJinrui4) Dive() string {
	return "ShinJinrui Dive!"
}

func InitDolphin() {
	aDolphin := &dolphin{}
	aSwhonJinrui4 := &ShinJinrui4{
		dolphin: aDolphin,
	}
	fmt.Println(aSwhonJinrui4.Dive())
	fmt.Println(aSwhonJinrui4.dolphin.Dive())
}
