// Last updated: 11/7/2025, 2:47:56 PM
func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }

    if len(s) + len(t) == 0 {
        return true
    }

    i := 0
    j := 0

    for i < len(t) {
        if j < len(s) && t[i] == s[j] {
            j++
        }
        if j == len(s) {
            return true
        }

        i++
    }   
    return false 
}