package main

func containsNearbyDuplicate(nums []int, k int) bool {
	indexesMap := make(map[int]int)

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for i := range nums {
		if j, exist := indexesMap[nums[i]]; exist {
			if abs(i - j) <= k {
				return true
			}
		}

		indexesMap[nums[i]] = i
	}

	return false
}