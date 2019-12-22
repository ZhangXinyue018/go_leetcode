package chip

type Chips interface {
	GetCount() int

	Test(i, j int) int

	IsGood(i int) bool
}

// chip issue, good or bad
func IdentityChip(chips Chips) int {
	testIndexes := make([]int, 0)
	n := chips.GetCount()
	for i := 0; i < n; i++ {
		testIndexes = append(testIndexes, i)
	}
	return IdentityChipSub(chips, testIndexes)
}

func IdentityChipSub(chips Chips, testIndexes []int) int {
	count := len(testIndexes)
	switch {
	case count == 1 || count == 2:
		return testIndexes[0]
	case count == 3:
		if chips.Test(testIndexes[0], testIndexes[1]) == 3 {
			return testIndexes[0]
		}
		return testIndexes[2]
	default:
		newTestIndexes := make([]int, 0)
		for i := 0; i < count-1; i = i + 2 {
			if chips.Test(testIndexes[i], testIndexes[i+1]) == 3 {
				newTestIndexes = append(newTestIndexes, i)
			}
		}
		if count%2 != 0 {
			goodCount := 0
			for i := 0; i < count-1; i++ {
				testResult := chips.Test(testIndexes[count-1], testIndexes[i])
				if testResult == 3 || testResult == 1 {
					goodCount++
				}
			}
			if goodCount >= (count-1)/2 {
				newTestIndexes = append(newTestIndexes, count-1)
			}
		}
		return IdentityChipSub(chips, newTestIndexes)
	}
}
