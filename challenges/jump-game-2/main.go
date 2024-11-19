package main

func jump(nums []int) int {
	count := len(nums) - 1
	if count < 1 {
		return 0
	}

	jumpCount, currentEnd, farthest := 0, 0, 0
	for i := 0; i <= count; i++ {
		farthest = max(farthest, i + nums[i])

		if i == currentEnd {
			jumpCount++
			currentEnd = farthest

			if currentEnd >= count {
				break
			}
		}
	} 

	return jumpCount
}