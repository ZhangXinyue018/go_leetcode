package stack

import (
	"go_leetcode/basic/commons"
	"testing"
)

func TestStack_General(t *testing.T) {
	t.Run("general test", func(t *testing.T) {
		stack := &Stack{}
		stack.Push(1)
		stack.Push(2)
		if !(commons.CheckArrayEqual(stack.Items, []int{1, 2})) {
			t.Errorf("[Push] Expect %v but get %v", []int{1, 2}, stack.Items)
		}
		stack.Push(6)
		stack.Push(4)
		stack.Push(9)
		popResult := stack.Pop()
		if popResult != 9 {
			t.Errorf("[Pop] Expect %v but get %v", 9, popResult)
		}
		if !(commons.CheckArrayEqual(stack.Items, []int{1, 2, 6, 4})) {
			t.Errorf("[Pop] Expect %v but get %v", []int{1, 2, 6, 4}, stack.Items)
		}
	})

}
