package main

import "strconv"

func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }

	var result []string
    start := 0

	for end := 1; end <= len(nums); end++ {
        if end == len(nums) || nums[end] > nums[end-1]+1 {
            if start == end-1 {
                result = append(result, strconv.Itoa(nums[start]))
            } else {
                result = append(result, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end-1]))
            }
            start = end
        }
    }

	return result
}