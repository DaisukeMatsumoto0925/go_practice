package handler

import (
	"math/rand"
	"sync"
	"time"

	"pancake.maker/xfpng345/api"
)

func init() {
	rand.Seed(time.Now(), UnixNano())
}

type BakerHandler struct {
	report *report
}

type report struct {
	sync.Mutex
	data map[api.Pancake_menu]int
}

func NewBakerHandler() *BakerHandler{
	return &BakerHandler{
		report: &report{
			data: make(map[api.Pancake_Menu]int)
		}
	}
}

func (h *BakerHandler) Bake(
	ctx context.Context,
	req *api.BakeRequest,
) (*api.BakeResponse,error){
	if req.Menu == api.Pancake_UNKNOWN
	|| req.Menu > api.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes,InvalidArgument, "パンケーキを選んでください！")
	}
}