package binary_heap

import (
	"go_leetcode/basic/commons"
	"reflect"
	"testing"
)

var (
	binaryHeapComparatorInt = BinaryHeapCompare{
		Less: func(a interface{}, b interface{}) bool {
			va := reflect.ValueOf(a)
			vb := reflect.ValueOf(b)
			return va.Int() < vb.Int()

		},
		Greater: func(a interface{}, b interface{}) bool {
			va := reflect.ValueOf(a)
			vb := reflect.ValueOf(b)
			return va.Int() > vb.Int()
		},
	}
)

func TestBinaryHeap_Insert(t *testing.T) {
	t.Run("test insert", func(t *testing.T) {
		binaryHeap := &BinaryHeap{Comparator: binaryHeapComparatorInt}
		type checkMap struct {
			input int
			check []int
		}
		maps := []checkMap{
			{5, []int{5}},
			{8, []int{8, 5}},
			{9, []int{9, 5, 8}},
			{10, []int{10, 9, 8, 5}},
			{15, []int{15, 10, 8, 5, 9}},
			{17, []int{17, 10, 15, 5, 9, 8}},
			{18, []int{18, 10, 17, 5, 9, 8, 15}},
			{6, []int{18, 10, 17, 6, 9, 8, 15, 5}},
		}

		for _, value := range maps {
			binaryHeap.Insert(value.input)
			if !commons.CheckArrayEqual(binaryHeap.Items, value.check) {
				t.Errorf("InsertSort() = %v, want %v", binaryHeap.Items, value.check)
			}
		}
	})
}

func TestBinaryHeap_Delete(t *testing.T) {
	t.Run("test delete", func(t *testing.T) {
		binaryHeap := &BinaryHeap{Comparator: binaryHeapComparatorInt}
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
