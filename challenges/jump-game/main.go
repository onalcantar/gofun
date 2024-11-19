package main

func canJump(nums []int) bool {
	count := len(nums)
	var position int

	for i := 0; i < count; i++ {
		if i > position {
			return false
		}

		position = max(position, i + nums[i])
		if position >= count - 1 {
			return true
		}
	}

	return true
}