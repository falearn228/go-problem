// Last updated: 11/7/2025, 2:47:06 PM
func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
        return ""
    }
    maxSubStrLen := gcd(len(str1), len(str2))

    return str1[:maxSubStrLen]
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
