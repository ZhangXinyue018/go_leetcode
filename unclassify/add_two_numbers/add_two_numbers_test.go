package add_two_numbers

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,}}},
				&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4,}}}},
			"708",
		},
		{
			"test 2",
			args{&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9,}}},
				&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1}}}}},
			"7042",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddTwoNumbers(tt.args.l1, tt.args.l2)
			resultGot := strconv.Itoa(got.Val)
			for got.Next != nil{
				resultGot = resultGot + strconv.Itoa(got.Next.Val)
				got = got.Next
			}
			if !reflect.DeepEqual(resultGot, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", resultGot, tt.want)
			}
		})
	}
}
