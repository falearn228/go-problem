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

func sortString(str string) string {
    sptStr := strings.Split(str, "")
    sort.Strings(sptStr)
    return strings.Join(sptStr, "")
}