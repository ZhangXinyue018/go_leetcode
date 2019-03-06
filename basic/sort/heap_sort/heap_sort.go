package heap_sort

import "go_leetcode/basic/datastructure/binary_heap"

// 方法：形成二叉堆（最大堆）之后，不断的删除根节点，以按照顺序由大到小排列

func HeapSort(nums []int) []int {
	binaryHeap := binary_heap.BinaryHeap{}
	for i := 0; i < len(nums); i++ {
		binaryHeap.Insert(nums[i])
	}
	result := []int{}
	for len(binaryHeap.Items) > 0 {
		result = append(result, binaryHeap.Items[0])
		binaryHeap.Delete(0)
	}
	return result
}
