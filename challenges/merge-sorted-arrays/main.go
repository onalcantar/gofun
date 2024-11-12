package main

import "fmt"

func main() {
	// Test case 1: Example from the prompt
    nums1 := []int{1, 2, 3, 0, 0, 0}
    m, nums2, n := 3, []int{2, 5, 6}, 3
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)  // Expected output: [1, 2, 2, 3, 5, 6]

    // Test case 2: nums2 is empty
    nums1 = []int{1, 2, 3}
    m, nums2, n = 3, []int{}, 0
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)  // Expected output: [1, 2, 3]

    // Test case 3: nums1 is empty (only has 0s as placeholders)
    nums1 = []int{0}
    m, nums2, n = 0, []int{1}, 1
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)  // Expected output: [1]

    // Test case 4: Both arrays have elements all equal
    nums1 = []int{2, 2, 2, 0, 0, 0}
    m, nums2, n = 3, []int{2, 2, 2}, 3
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)  // Expected output: [2, 2, 2, 2, 2, 2]

    // Test case 5: Larger, mixed arrays
    nums1 = []int{1, 4, 7, 0, 0, 0}
    m, nums2, n = 3, []int{2, 5, 6}, 3
    merge(nums1, m, nums2, n)
    fmt.Println(nums1)  // Expected output: [1, 2, 4, 5, 6, 7]
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, k := m - 1, n - 1, m + n - 1

    for j >= 0 {
        if i >= 0 && nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
}