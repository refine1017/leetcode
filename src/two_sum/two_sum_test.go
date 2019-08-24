package two_sum

import (
	"testing"
	"reflect"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums []int
		target int
		output []int
	}{
		{[]int{2,7,11,15}, 9, []int{0,1}},
		{[]int{-3,0,3,15}, 0, []int{0,2}},
	}

	for _, tt := range tests {
		if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(tt.output, got) {
			t.Errorf("twoSum(%v, %v) = %v, want = %v", tt.nums, tt.target, got, tt.output)
		}
	}
}
