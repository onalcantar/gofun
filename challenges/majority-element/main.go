package main

func majorityElement(nums []int) int {
    var rptts = make(map[int]int)
	for _, num := range nums {
		rptts[num]++;
	}

	k, h := 0, 0
	for num, rptt := range rptts {
		if rptt > h {
			h, k = rptt, num
		}
	}

	return k
}