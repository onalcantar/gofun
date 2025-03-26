package main

import (
	"fmt"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		s           string
		t 			string
        expected   	bool
	}{
		{
			s: 			"abc",
			t: 			"ahbgdc",
			expected: 	true,
		},
		{
			s: 			"axc",
			t: 			"ahbgdc",
			expected: 	false,
		},
		{
			s: 			"abc",
			t: 			"ahbkmclapbtc",
			expected: 	true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%t", tt.expected), func(t *testing.T) {
            // Call the function
            k := isSubsequence(tt.s, tt.t)

            // Check if k is correct
			if (k != tt.expected) {
				t.Errorf("Expected %t, got %t", k, tt.expected)
			}
        })
	}
}
