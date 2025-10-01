func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	for i := range s {
		if len(stack) == 0 {
			stack = append(stack, s[i])
            continue
		}
		for len(stack) > 0 {
			if s[i] == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				break
			} else if s[i] == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				break
			} else if s[i] == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				break
			} else {
				stack = append(stack, s[i])
				break
			}
		}

	}
	return len(stack) == 0
}
