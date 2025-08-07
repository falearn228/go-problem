func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    // freq1 == freq2
    freqMap1 := make(map[rune]int)
    freqMap2 := make(map[rune]int) 
    for _, v := range word1 {
        freqMap1[v]++
    }
    for _, v := range word2 {
        freqMap2[v]++
        if _, ok := freqMap1[v]; !ok {
            return false
        }
    }

    freq1 := make([]int, len(freqMap1))
    ind := 0
    for _, v := range freqMap1 {
        freq1[ind] = v
        ind++
    }
    freq2 := make([]int, len(freqMap2))
    ind = 0
    for _, v := range freqMap2 {
        freq2[ind] = v
        ind++
    }

    slices.Sort(freq1)
    slices.Sort(freq2)

    if !slices.Equal(freq1, freq2) {
        return false
    }
    return true
}