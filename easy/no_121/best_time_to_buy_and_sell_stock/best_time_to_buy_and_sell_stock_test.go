package best_time_to_buy_and_sell_stock

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"test2", args{[]int{7, 6, 5, 4, 3, 1}}, 0},
		{"test3", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitClearer(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"test2", args{[]int{7, 6, 5, 4, 3, 1}}, 0},
		{"test3", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitClearer(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitClearer() = %v, want %v", got, tt.want)
			}
		})
	}
}
