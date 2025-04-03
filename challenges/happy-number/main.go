package main

func isHappy(n int) bool {
    trackedSums := make(map[int]bool)
    for n != 1 {
		sum := 0
        for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10 
		}

		n = sum 
		if n == 1 {
			break
		}

		if exist := trackedSums[n]; exist {
			break
		}

		trackedSums[n] = true
    }

	return n == 1
}