package main

func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
		return len(nums)
	}

	k := 1

	for _, num := range nums[2:] {
		if num != nums[k -1] {
			nums[k+1] = num
			k++
		}
	}

	return k + 1
}