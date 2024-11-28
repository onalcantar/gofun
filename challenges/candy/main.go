package main

func candy(ratings []int) int {
	count := len(ratings)
	candies := make([]int, count)
	
	for i := 1; i < count; i++ {
		if ratings[i] > ratings[i - 1] {
			candies[i] = candies[i - 1] + 1
		}
	}

	for i := count - 2; i >= 0; i-- {
		if ratings[i] > ratings[i + 1] {
			candies[i] = max(candies[i], candies[i + 1] + 1)
		}
	}

	return sum(candies) + count
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}