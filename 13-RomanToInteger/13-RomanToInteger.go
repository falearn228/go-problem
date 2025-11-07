// Last updated: 11/7/2025, 2:49:34 PM
func romanToInt(s string) int {
	input := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	output := []int{1, 5, 10, 50, 100, 500, 1000}

	ans := 0
	for i := range s {
		rimlan := slices.Index(input, s[i])
        num := output[rimlan]
        if i+1 < len(s) && ((s[i] == 'I') && (s[i+1] == 'V' || s[i+1] == 'X')) {
            num = -output[rimlan]
        } else if i+1 < len(s) && ((s[i] == 'X') && (s[i+1] == 'L' || s[i+1] == 'C')) {
            num = -output[rimlan]
        } else if i+1 < len(s) && ((s[i] == 'C') && (s[i+1] == 'D' || s[i+1] == 'M')) {
            num = -output[rimlan]
        }
		ans += num
	}

	return ans
}