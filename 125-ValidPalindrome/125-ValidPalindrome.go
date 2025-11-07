// Last updated: 11/7/2025, 2:48:53 PM
func isPalindrome(s string) bool {
    s = isLetter(s)
    s = strings.ToLower(s)

    if len(s) == 0 {
        return true
    }
    left, right := 0, len(s)-1

    for left < right {
        if s[left] != s[right] {
            return false
        }

        left++
        right--
    }

    return true
}

func isLetter(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			b.WriteRune(v)
		}
	}
	return b.String()
}
