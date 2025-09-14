func longestCommonPrefix(strs []string) string {
    minLen := 1_000
    for i := range strs {
        minLen = min(minLen, len(strs[i]))
    }

    var ans strings.Builder
    ans.Grow(minLen)
    for i := 0; i < minLen; i++ {
        pref := strs[0][i]
        for j := 1; j < len(strs); j++ {
            fmt.Println(strs[j][i], pref)
            if strs[j][i] != pref {
                return ans.String()
            }  
        }
        ans.WriteByte(pref) 
    }

    return ans.String()
}