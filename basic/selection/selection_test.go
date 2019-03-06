package selection

import "testing"

func TestSelection(t *testing.T) {
	type args struct {
		nums      []int
		targetInt int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{6, 7, 3, 1, 4, 2, 5}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Selection(tt.args.nums, tt.args.targetInt); got != tt.want {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}
