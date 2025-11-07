// Last updated: 11/7/2025, 2:48:05 PM
func wordPattern(pattern string, s string) bool {
    govno := make(map[string]byte)
    mocha := make(map[byte]string)

    fields := strings.Fields(s)
    if len(fields) != len(pattern) {
        return false
    }
    for i := range fields {
        char1, ok1 := govno[fields[i]]
        char2, ok2 := mocha[pattern[i]]
        if ok1 {
            if char1 != pattern[i] {
                return false
            }
        }

        if ok2 {
            if char2 != fields[i] {
                return false
            }
        }
        govno[fields[i]] = pattern[i]
        mocha[pattern[i]] = fields[i]
    }

    return true
}