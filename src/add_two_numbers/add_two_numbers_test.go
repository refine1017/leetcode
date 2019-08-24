package add_two_numbers

import (
	"testing"
	"reflect"
)

func create(vals []int) *ListNode {
	root := &ListNode{}
	current := root
	for i:=0; i<len(vals);i++ {
		current.Val = vals[i]
		if i != len(vals) - 1 {
			current.Next = &ListNode{}
			current = current.Next
		}
	}
	return root
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1 *ListNode
		l2 *ListNode
		output *ListNode
	}{
		{create([]int{2,4,3}), create([]int{5,6,4}), create([]int{7,0,8})},
		{create([]int{9}), create([]int{9}), create([]int{8,1})},
	}

	for _, tt := range tests {
		if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(tt.output, got) {
			t.Errorf("twoSum(%v, %v) = %+v, want = %+v", tt.l1, tt.l2, got, tt.output)
		}
	}
}
