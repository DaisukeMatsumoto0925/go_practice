package status

import "fmt"

type Status int

const (
	Good Status = iota
	Tired
)

func (s Status) String() string {
	switch s {
	case Good:
		return "Good!"
	case Tired:
		return "Tired..."
	default:
		return ""
	}
}

type poorGrasshopper struct {
	status Status
}

func (g *poorGrasshopper) HighJUmp() {
	fmt.Println("High Jump!")
	g.status = Tired
}

type ShinJinrui3 struct{
	status Status
	*poorGrasshopper
}

func InitStatus() {
	aPoorGrasshoper := &poorGrasshopper{
		status: Good,
	}
	aShinJinrui3 := &ShinJinrui3{
		status:  Good,
		poorGrasshopper: aPoorGrasshoper,
	}
	aShinJinrui3.HighJUmp()
	fmt.Println("aPoorGrasshopper is", aPoorGrasshoper.status)
	fmt.Println("aShinJinrui3 is", aShinJinrui3.status)
}
