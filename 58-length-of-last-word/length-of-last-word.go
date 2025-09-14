func lengthOfLastWord(s string) int {
    ans := strings.Fields(s)
    return len(ans[len(ans)-1])
}