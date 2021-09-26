package bisect_test

import (
	"testing"

	"github.com/DaisukeMatsumoto0925/go-test/bisect"
)

func TestStop_block(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 4

	ans := bisect.BisectLeft(a, key)
	if ans != 3 {
		t.Errorf("get wrong ans: %d", ans)
	}
}
