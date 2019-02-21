package min_stack

import (
	"fmt"
	"testing"
)

func TestGeneral(t *testing.T) {
	test := Constructor()
	test.Push(0)
	test.Push(1)
	test.Push(0)
	fmt.Println(test.GetMin())
	test.Pop()
	fmt.Println(test.GetMin())
}
