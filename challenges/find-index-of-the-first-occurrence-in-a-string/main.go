package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen < needleLen {
		return -1
	}

	for i := 0; i <= haystackLen - needleLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	
	return -1
}