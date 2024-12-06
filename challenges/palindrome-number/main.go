package main

import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}

	numStr := strconv.Itoa(x)
	numReverse := reverseStr(numStr)

	return numStr == numReverse
}

func reverseStr(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}