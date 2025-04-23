package main

func maxArea(height []int) int {
    if (len(height) == 0) {
		return 0
	}

	left, right := 0, len(height) - 1
	maxArea := 0

	for left < right {
		currentWidth := right - left

		heightLeft := height[left]
		heightRight := height[right]
		currentHeight := min(heightLeft, heightRight)

		maxArea = max(maxArea, currentWidth * currentHeight)

		if heightLeft <= heightRight {
			left++
		} else {
			right--
		}
	}

	return maxArea
}