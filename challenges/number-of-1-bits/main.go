package main

func hammingWeight(num int) int {
    count := 0

    for range 32 {
		count += num & 1
        num = num >> 1
    }
	
    return count
}