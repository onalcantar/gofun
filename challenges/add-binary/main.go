package main

func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }

    indexB := len(b) - 1
    res := make([]byte, len(a))

    var sum, shift byte
    
    for i := len(a) - 1; i >= 0; i-- {
        sum = shift + a[i]
        if indexB >= 0 {
            sum += b[indexB]
            indexB--
        }

        res[i] = sum % 2 + '0'
        
        shift = sum >> 1 % 2
    }

    if shift == 0 {
		return string(res)
    }

	return "1" + string(res)
}