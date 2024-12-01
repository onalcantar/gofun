package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var prefix string
	firstStr := strs[0]

	for i := range firstStr {
		for _, str := range strs[1:] {
			if i == len(str) || str[i] != firstStr[i] {
				return prefix
			}
		}

		prefix += fmt.Sprintf("%c", firstStr[i])
	}

	return prefix
}

