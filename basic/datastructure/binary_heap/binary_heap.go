package binary_heap

// 最大堆

type BinaryHeap struct {
	Items      []interface{}
	Comparator BinaryHeapCompare
}

type BinaryHeapCompare struct {
	Less func(a interface{}, b interface{}) bool

	Greater func(a interface{}, b interface{}) bool
}

func (binaryHeap *BinaryHeap) Insert(num interface{}) {
	binaryHeap.Items = append(binaryHeap.Items, num)
	i := len(binaryHeap.Items) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if parent >= 0 && binaryHeap.Comparator.Less(binaryHeap.Items[parent], binaryHeap.Items[i]) {
			binaryHeap.Items[parent], binaryHeap.Items[i] = binaryHeap.Items[i], binaryHeap.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (binaryHeap *BinaryHeap) Delete(index int) {
	binaryHeap.Items[index] = binaryHeap.Items[len(binaryHeap.Items)-1]
	binaryHeap.Items = binaryHeap.Items[:len(binaryHeap.Items)-1]
	i := index
	for i < len(binaryHeap.Items) {
		child1 := 2*i + 1
		child2 := 2*i + 2
		var target int
		if child1 < len(binaryHeap.Items) && child2 < len(binaryHeap.Items) {
			if binaryHeap.Comparator.Greater(binaryHeap.Items[child1], binaryHeap.Items[child2]) {
				target = child1
			} else {
				target = child2
			}
		} else if child1 < len(binaryHeap.Items) {
			target = child1
		} else if child2 < len(binaryHeap.Items) {
			target = child2
		} else {
			break
		}

		if binaryHeap.Comparator.Less(binaryHeap.Items[i], binaryHeap.Items[target]) {
			binaryHeap.Items[i], binaryHeap.Items[target] = binaryHeap.Items[target], binaryHeap.Items[i]
			i = target
		} else {
			break
		}
	}
}
