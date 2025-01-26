package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		x           string
        expected   	bool
	}{
		{
			x: 			"A man, a plan, a canal: Panama",
			expected: 	true,
		},
		{
			x: 			"race a car",
			expected: 	false,
		},
		{
			x: 			" ",
			expected: 	true,
		},
		{
			x: 			"0P",
			expected: 	false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%t", tt.expected), func(t *testing.T) {
            // Call the function
            k := isPalindrome(tt.x)

            // Check if k is correct
			if (k != tt.expected) {
				t.Errorf("Expected %t, got %t", k, tt.expected)
			}
        })
	}
}
