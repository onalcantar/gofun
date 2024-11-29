package main

var romansMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
    total, num, prev_num := 0, 0, 0

	for i := 0; i < len(s); i++ {
		prev_num = num
		num = romansMap[s[i]]

		if num > prev_num {
			total = total + num - 2 * prev_num
		} else {
			total = total + num
		}
	}

	return total
}