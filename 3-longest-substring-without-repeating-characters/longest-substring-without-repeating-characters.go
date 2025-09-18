func lengthOfLongestSubstring(s string) int {
    length := 0
    currLeng := 0
    hash := make(map[byte]int)

    i := 0
    for ; i < len(s); i++ {
        index, ok := hash[s[i]]
        if !ok {
            currLeng++
            hash[s[i]] = i
            length = max(length, currLeng)
        } else {
            currLeng = 0
            hash = nil
            hash = make(map[byte]int)
            i = index    
        }
    }

    return length
}