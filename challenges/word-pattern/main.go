package main

const Space = 32 // Space in bytes

func wordPattern(pattern string, s string) bool {	
	var words []string
	var word []rune
	wordsPattern := make(map[byte]string)
	uniqueWords := make(map[string]bool)

	for i, c := range s {
		if c != Space  {
			word = append(word, c)
		} 
		
		if c == Space || i == len(s) - 1 {
			words = append(words, string(word))
			word = nil
		}
	}

	if len(words) != len(pattern) {
		return false
	}

	for i, w := range words {
		chartAt := pattern[i]
		if _, ok := wordsPattern[chartAt]; !ok && !uniqueWords[w] {
			wordsPattern[chartAt] = w
			uniqueWords[w] = true
		}

		if wordsPattern[chartAt] != w {
			return false
		}
	}

	return true;
}