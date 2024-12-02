package main

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack  	string
		needle 		string
		expected 	int
	}{
		{
			haystack: 	"sadbutsad",
			needle: 	"sad",
			expected: 	0,
		},
		{
			haystack: 	"leetcode",
			needle: 	"leeto",
			expected: 	-1,
		},
		{
			haystack: 	"leetcodeleeto",
			needle: 	"leeto",
			expected: 	8,
		},
		{
			haystack: 	"",
			needle: 	"",
			expected: 	0,
		},
		{
			haystack: 	"foobar",
			needle: 	"bar",
			expected: 	3,
		},
		{
			haystack: 	"foobar",
			needle: 	"bart",
			expected: 	-1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testing with expected=%d", tt.expected), func(t *testing.T) {
			// Call the function
			k := strStr(tt.haystack, tt.needle)

			// Check if k is correct
			if k != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, k)
			}
		})
	}
}
