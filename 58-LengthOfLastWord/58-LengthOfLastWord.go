// Last updated: 11/7/2025, 2:49:08 PM
func lengthOfLastWord(s string) int {
    ans := 0
    isFirstSpace := true

    for i := len(s)-1; i >= 0; i-- {
        if isFirstSpace && s[i] == ' ' {
            continue
        } else {
            isFirstSpace = false
        }

        if s[i] != ' ' {
            ans += 1
        } else {
            break
        }
    }

    return ans
}