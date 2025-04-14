package main

func isValid(s string) bool {
	stack := []rune{}
    hash := map[rune]rune{
		')': '(', 
		']': '[', 
		'}': '{',
	}

	for _, c := range s {
		if match, exists := hash[c]; exists {
			if len(stack) > 0 && stack[len(stack) - 1] == match {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}

			continue
		}

		stack = append(stack, c)
	}

	return len(stack) == 0
}