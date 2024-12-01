package main

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs  		[]string
		expected 	string
	}{
		{
			strs: 		[]string{"flower","flow","flight"},
			expected: 	"fl",
		},
		{
			strs: 		[]string{"dog","racecar","car"},
			expected: 	"",
		},
		{
			strs: 		[]string{"dog"},
			expected: 	"dog",
		},
		{
			strs: 		[]string{"", ""},
			expected: 	"",
		},
		{
			strs: 		[]string{"ab", "a"},
			expected: 	"a",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%v", tt.expected), func(t *testing.T) {
			// Call the function
			k := longestCommonPrefix(tt.strs)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, k)
			}
		})
	}
}
