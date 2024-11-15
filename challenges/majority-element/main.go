package main

func majorityElement(nums []int) int {
    rptts := make(map[int]int)
    n := len(nums)

    for _, num := range nums {
        rptts[num]++
        if rptts[num] > n / 2 {
            return num
        }
    }

    return -1
}