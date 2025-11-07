// Last updated: 11/7/2025, 2:47:55 PM
func decodeString(s string) string {

	stack := make([]byte, 0)
	for i := range s {
		if s[i] == ']' {
			tmp := make([]byte, 0)
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			}
			for l, r := 0, len(tmp)-1; l < r; l, r = l+1, r-1 {
				tmp[l], tmp[r] = tmp[r], tmp[l]
			}
			digits := make([]byte, 0)
			for len(stack) > 0 {
				b := stack[len(stack)-1]
				if b < '0' || b > '9' {
					break
				}
				digits = append(digits, b)
				stack = stack[:len(stack)-1]
			}
			for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
				digits[l], digits[r] = digits[r], digits[l]
			}
			k := 0
			if len(digits) > 0 {
				k, _ = strconv.Atoi(string(digits))
			}
			for range k {
				stack = append(stack, tmp...)
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}