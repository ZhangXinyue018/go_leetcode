package min_stack

import (
	"fmt"
	"testing"
)

func TestConstructorBetter(t *testing.T) {
	test := ConstructorBetter()
	test.PushBetter(2147483646)
	test.PushBetter(2147483646)
	test.PushBetter(2147483647)
	fmt.Println(test.TopBetter())
	test.PopBetter()
	fmt.Println(test.GetMinBetter())
	test.PopBetter()
	fmt.Println(test.GetMinBetter())
	test.PopBetter()
	test.PushBetter(2147483647)
	fmt.Println(test.TopBetter())
	fmt.Println(test.GetMinBetter())
	test.PushBetter(2147483648)
	fmt.Println(test.TopBetter())
	fmt.Println(test.GetMinBetter())
	test.PopBetter()
	fmt.Println(test.GetMinBetter())
}

func TestConstructorBetter2(t *testing.T) {
	test := ConstructorBetter()
	test.PushBetter(-2)
	test.PushBetter(0)
	test.PushBetter(-3)
	fmt.Println(test.GetMinBetter())
	test.PopBetter()
	fmt.Println(test.TopBetter())
	fmt.Println(test.GetMinBetter())
}