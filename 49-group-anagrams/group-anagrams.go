func groupAnagrams(strs []string) [][]string {
    anagrams := make(map[string][]string)
    for i := range strs {
        str := sortString(strs[i])
        anagrams[str] = append(anagrams[str], strs[i])
    }

    res := make([][]string, 0, len(anagrams))
    for i := range anagrams {
        res = append(res, anagrams[i])
    }

    return res
}

func sortString(s string) string {
    bytes := []byte(s)
    sort.Slice(bytes, func(i, j int) bool {
        return bytes[i] < bytes[j]
    })
    return string(bytes)
}