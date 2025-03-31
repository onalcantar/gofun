package main

func twoSum(nums []int, target int) []int {
	numsPositionMap := make(map[int]int)

	for i, v := range nums {
		if k, ok := numsPositionMap[target-v]; ok {
			return []int{i, k}
		}

		numsPositionMap[v] = i
	}

	return []int{}
}
