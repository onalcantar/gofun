package main

func removeDuplicates(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}
    
	k := 1
    
	for i, num := range nums[1:] {
		if (num != nums[i]) {
			nums[k] = num
			k++
		}
	}

	return k
}