package chip

import (
	"math/rand"
	"testing"
)

func TestIdentityChip(t *testing.T) {
	testChips := &ChipsImplementation{ChipItems: []int{1, 0, 0, 1, 1, 1}}
	//testChips := &ChipsImplementation{ChipItems: []int{1, 0, 0, 1, 0, 1, 1}}
	result := IdentityChip(testChips)
	if !testChips.IsGood(result) {
		t.Errorf("result failed %v\n", result)
	}
}

type ChipsImplementation struct {
	ChipItems []int
}

func (impl *ChipsImplementation) GetCount() int {
	return len(impl.ChipItems)
}

func (impl *ChipsImplementation) Test(i, j int) int {
	if impl.ChipItems[i] == 1 && impl.ChipItems[j] == 1 {
		return 3
	} else if impl.ChipItems[i] == 0 && impl.ChipItems[j] == 0 {
		return rand.Int() % 4
	} else if impl.ChipItems[i] == 1 {
		return rand.Int() % 2
	} else {
		return (rand.Int() % 2) * 2
	}
}

func (impl *ChipsImplementation) IsGood(i int) bool {
	return impl.ChipItems[i] == 1
}
