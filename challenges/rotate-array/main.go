package main

func rotate(nums []int, k int) {
	total := len(nums)
	if total < 2 || k < 1 {
		return
	}

	k %= total

	reverse(nums, 0, total - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, total - 1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}