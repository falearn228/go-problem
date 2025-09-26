func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    charCount := make(map[byte]int)
    for i := range s {
        charCount[s[i]]++
    }

    for i := range t {
        cnt, ok := charCount[t[i]]
        if !ok || cnt < 1 {
            return false
        }
        charCount[t[i]]--
    }
    return true
}