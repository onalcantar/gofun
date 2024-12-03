package main

import (
	"fmt"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		s  			string
		expected 	string
	}{
		{
			s: 			"the sky is blue",
			expected: 	"blue is sky the",
		},
		{
			s: 			"  hello world  ",
			expected: 	"world hello",
		},
		{
			s: 			"a good   example",
			expected: 	"example good a",
		},
		{
			s: 			" good ",
			expected: 	"good",
		},
		{
			s: 			"hello",
			expected: 	"hello",
		},
		{
			s: 			"                   hello",
			expected: 	"hello",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%s", tt.expected), func(t *testing.T) {
			// Call the function
			k := reverseWords(tt.s)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, k)
			}
		})
	}
}
