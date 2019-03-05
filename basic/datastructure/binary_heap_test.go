package datastructure

import (
	"go_leetcode/basic/commons"
	"testing"
)

func TestBinaryHeap_Insert(t *testing.T) {
	t.Run("test insert", func(t *testing.T) {
		binaryHeap := &BinaryHeap{}
		binaryHeap.Insert(5)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{5})
		}
		binaryHeap.Insert(8)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{8, 5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{8, 5})
		}
		binaryHeap.Insert(9)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{9, 5, 8}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{9, 5, 8})
		}
		binaryHeap.Insert(10)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{10, 9, 8, 5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{10, 9, 8, 5})
		}
		binaryHeap.Insert(15)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{15, 10, 8, 5, 9}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{15, 10, 8, 5, 9})
		}
		binaryHeap.Insert(17)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{17, 10, 15, 5, 9, 8}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{17, 10, 15, 5, 9, 8})
		}
		binaryHeap.Insert(18)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{18, 10, 17, 5, 9, 8, 15}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{18, 10, 17, 5, 9, 8, 15})
		}
		binaryHeap.Insert(6)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{18, 10, 17, 6, 9, 8, 15, 5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{18, 10, 17, 6, 9, 8, 15, 5})
		}
	})
}

func TestBinaryHeap_Delete(t *testing.T) {
	t.Run("test delete", func(t *testing.T) {
		binaryHeap := &BinaryHeap{}
		binaryHeap.Insert(5)
		binaryHeap.Insert(8)
		binaryHeap.Insert(9)
		binaryHeap.Insert(10)
		binaryHeap.Insert(15)
		binaryHeap.Insert(17)
		binaryHeap.Insert(18)
		binaryHeap.Insert(6)
		binaryHeap.Delete(2)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{18, 10, 15, 6, 9, 8, 5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{18, 10, 15, 6, 9, 8, 5})
		}
		binaryHeap.Delete(0)
		if !commons.CheckArrayEqual(binaryHeap.Items, []int{15, 10, 8, 6, 9, 5}) {
			t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, []int{15, 10, 8, 6, 9, 5})
		}
		//for len(binaryHeap.Items) > 0 {
		//	fmt.Println("=============")
		//	fmt.Printf("before delete current len %d \n", len(binaryHeap.Items))
		//	binaryHeap.Delete(0)
		//	fmt.Printf("after delete current len %d \n", len(binaryHeap.Items))
		//}
	})
}
