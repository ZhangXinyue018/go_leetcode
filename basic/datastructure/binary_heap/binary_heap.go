package binary_heap

// 最大堆

type BinaryHeap struct {
	Items []int
}

func (binaryHeap *BinaryHeap) Insert(num int) {
	binaryHeap.Items = append(binaryHeap.Items, num)
	i := len(binaryHeap.Items) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if parent >= 0 && binaryHeap.Items[parent] < binaryHeap.Items[i] {
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
			if binaryHeap.Items[child1] > binaryHeap.Items[child2] {
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

		if binaryHeap.Items[i] < binaryHeap.Items[target] {
			binaryHeap.Items[i], binaryHeap.Items[target] = binaryHeap.Items[target], binaryHeap.Items[i]
			i = target
		} else {
			break
		}
	}
}
