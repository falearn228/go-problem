func removeStars(s string) string {
    stack := make([]byte, 0, len(s))

    for i := range s {
        if s[i] == '*' && len(stack) > 0 {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }

    var sb strings.Builder
    sb.Grow(len(stack))
    for i := range stack {
        sb.WriteByte(stack[i])
    }

    return sb.String()
}