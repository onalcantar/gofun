package main

import(
	"fmt"
	"testing"
)

func Test_convert(t *testing.T) {
	tests := []struct {
		s  			string
		numRows 	int
		expected 	string
	}{
		{
			s: 			"PAYPALISHIRING",
			numRows: 	3,
			expected: 	"PAHNAPLSIIGYIR",
		},
		{
			s: 			"PAYPALISHIRING",
			numRows: 	4,
			expected: 	"PINALSIGYAHRPI",
		},
		{
			s: 			"PAYPALISHIRING",
			numRows: 	5,
			expected: 	"PHASIYIRPLIGAN",
		},
		{
			s: 			"ABC",
			numRows: 	2,
			expected: 	"ACB",
		},
		{
			s: 			"A",
			numRows: 	1,
			expected: 	"A",
		},
		{
			s: 			"A",
			numRows: 	5,
			expected: 	"A",
		},
		{
			s: 			"AB",
			numRows: 	2,
			expected: 	"AB",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%s", tt.expected), func(t *testing.T) {
			// Call the function
			k := convert(tt.s, tt.numRows)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, k)
			}
		})
	}
}
