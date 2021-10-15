package leetcode

import (
	"fmt"
	"testing"
)

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a []int
		b int
		want []int
	}{
		{[]int{2,7,11,15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := twoSum(tt.a, tt.b)
			if ans[0] != tt.want[0] || ans[1] != tt.want[1] {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}