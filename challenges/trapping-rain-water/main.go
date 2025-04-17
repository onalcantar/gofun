package main

func trap(height []int) int {
    if (len(height) == 0) {
		return 0
	}

	left, right := 0, len(height) - 1
	maxLeft, maxRight := height[left], height[right]
	totalWater := 0

	for left < right {
		if maxLeft < maxRight {
			left++
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				totalWater += maxLeft - height[left]
			}
		} else {
			right--
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				totalWater += maxRight - height[right]
			}
		}
	}

	return totalWater
}