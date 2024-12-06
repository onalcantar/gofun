package main

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		x           int
        expected   	bool
	}{
		{
			x: 			121,
			expected: 	true,
		},
		{
			x: 			-121,
			expected: 	false,
		},
		{
			x: 			10,
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
