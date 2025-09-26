func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }

    anagrams := make(map[string][]string)
    for _, s := range strs {
        freq := make([]int, 26)
        for _ ,ch := range s {
            freq[ch-'a']++
        }

        var keyBd strings.Builder
        for _, count := range freq {
            keyBd.WriteByte('#')
            keyBd.WriteString(strconv.Itoa(count))
        } 
        key := keyBd.String()
        anagrams[key] = append(anagrams[key], s)
    }

    res := make([][]string, 0, len(anagrams))
    for _, group := range anagrams {
        res = append(res, group)
    }
    return res
}
