package main

import (
	"fmt"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		num  		int
		expected 	string
	}{
		{
			num: 		1,
			expected: 	"I",
		},
		{
			num: 		3,
			expected: 	"III",
		},
		{
			num: 		4,
			expected: 	"IV",
		},
		{
			num: 		5,
			expected: 	"V",
		},
		{
			num: 		10,
			expected: 	"X",
		},
		{
			num: 		40,
			expected: 	"XL",
		},
		{
			num: 		3749,
			expected: 	"MMMDCCXLIX",
		},
		{
			num: 		58,
			expected: 	"LVIII",
		},
		{
			num: 		1994,
			expected: 	"MCMXCIV",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%s", tt.expected), func(t *testing.T) {
			// Call the function
			k := intToRoman(tt.num)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, k)
			}
		})
	}
}
