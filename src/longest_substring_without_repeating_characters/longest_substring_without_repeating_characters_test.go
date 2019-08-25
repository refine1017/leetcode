package longest_substring_without_repeating_characters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s string
		output int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, tt := range tests {
		if got := lengthOfLongestSubstring(tt.s); tt.output != got {
			t.Errorf("twoSum(%v) = %v, want = %v", tt.s, got, tt.output)
		}
	}
}
