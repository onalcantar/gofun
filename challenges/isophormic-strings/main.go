package main

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	for i := range s {
		c1, c2 := s[i], t[i]

		if v, ok := mapST[c1]; ok && v != c2 {
			return false
		}

		if v, ok := mapTS[c2]; ok && v != c1  {
			return false
		}

		mapST[c1] = c2
		mapTS[c2] = c1
	}

	return true
}