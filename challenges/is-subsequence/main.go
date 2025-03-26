package main

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
    if m > n {
		return false
	}

	i := 0
	for j := 0; j < n && i < m; j++ {
		if s[i] == t[j] {
			i++
		}
    }

	return i == m;
}