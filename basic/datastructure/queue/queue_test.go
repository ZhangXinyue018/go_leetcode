package queue

import (
	"go_leetcode/basic/commons"
	"testing"
)

func TestQueue_General(t *testing.T) {
	t.Run("test general", func(t *testing.T) {
		queue := &Queue{}
		queue.Enqueue(6)
		queue.Enqueue(9)
		queue.Enqueue(1)
		queue.Enqueue(8)
		if !commons.CheckArrayEqual(queue.Items, []int{6, 9, 1, 8}) {
			t.Errorf("Queue.Enqueue() = %v, want %v", queue.Items, []int{6, 9, 1, 8})
		}
		deQResult := queue.Dequeue()
		if deQResult != 6 {
			t.Errorf("Queue.Dequeue() = %v, want %v", deQResult, 6)
		}
		if !commons.CheckArrayEqual(queue.Items, []int{9, 1, 8}) {
			t.Errorf("Queue.Dequeue = %v, want %v", queue.Items, []int{9, 1, 8})
		}
	})

}
