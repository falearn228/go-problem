func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

    pairs := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for i := range s {
        char := s[i]
        
        if open, isClosing := pairs[char]; isClosing {
            if len(stack) == 0 || stack[len(stack)-1] != open {
               return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, char)
        }
    }


	return len(stack) == 0
}
