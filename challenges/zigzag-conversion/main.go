package main

import "strings"

func convert(s string, numRows int) string {
	strLen := len(s)
    if numRows < 2 || strLen <= numRows {
		return s
	}

	var result strings.Builder
	rowOffset := (2 * numRows - 2)

	for row := range numRows {
		for i := range(strLen) {
			if (i - row) % rowOffset == 0 || (i + row) % rowOffset == 0 {
				result.WriteByte(s[i])
			}
		}
	}

	return result.String()
}