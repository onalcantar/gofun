package main

func productExceptSelf(nums []int) []int {
    count := len(nums)
    result := make([]int, count)

    prefix := 1
    for i := 0; i < count; i++ {
        result[i], prefix = prefix, prefix * nums[i]
    }

    suffix := 1
    for i := count - 1; i >= 0; i-- {
        result[i], suffix = result[i] * suffix, suffix * nums[i]
    }

    return result
}