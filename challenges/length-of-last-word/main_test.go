package main

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		s  			string
		expected 	int
	}{
		{
			s: 			"Hello World",
			expected: 	5,
		},
		{
			s: 			"   fly me   to   the moon  ",
			expected: 	4,
		},
		{
			s: 			"luffy is still joyboy",
			expected: 	6,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
			// Call the function
			k := lengthOfLastWord(tt.s)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, k)
			}
		})
	}
}
