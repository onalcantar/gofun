package main

func climbStairs(n int) int {
    beforeLast, last := 0, 1

    for i := 1; i <= n; i++ {
        beforeLast, last = last, beforeLast + last
    }

    return last
}