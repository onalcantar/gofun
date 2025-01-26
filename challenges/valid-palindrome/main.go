package main

func isPalindrome(s string) bool {
    if (len(s) == 0) {
        return true
    }

    left, right := 0, len(s)-1
    for left < right {
        for left < right && !isAlphaNumeric(s[left]) {
            left++
        }
        
        for left < right && !isAlphaNumeric(s[right]) {
            right--
        }

        if toLower(s[left]) != toLower(s[right]) {
            return false
        }

        left++
        right--
    }

    return true
}

func isAlphaNumeric(char byte) bool {
    return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z' || (char >= '0' && char <= '9'))
}

func toLower(char byte) byte {
    if char >= 'A' && char <= 'Z' {
        return char + ('a' - 'A')
    }

    return char
}