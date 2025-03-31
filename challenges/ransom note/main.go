package main

const CharA = 'a'

func canConstruct(ransomNote string, magazine string) bool {
    counters := [26]int{}

	for _, c := range magazine {
		counters[c - CharA]++
	}

	for _, c := range ransomNote {
		counters[c - CharA]--
		if counters[c - CharA] < 0 {
			return false
		}
	}

	return true
}