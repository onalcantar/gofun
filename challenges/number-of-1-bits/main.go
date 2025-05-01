package main

func hammingWeight(num int) int {
    count := 0

    for range 32 {
        if num&1 == 1 {
            count++
        }

        num = num >> 1
    }
	
    return count
}