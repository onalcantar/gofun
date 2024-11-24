package main

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	h := 0
	for _, citation := range citations {
		if citation < h + 1 {
			break
		}

		h++
	}

	return h
}