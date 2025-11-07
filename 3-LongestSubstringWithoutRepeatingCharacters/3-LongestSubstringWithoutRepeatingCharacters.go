// Last updated: 11/7/2025, 2:49:45 PM
func lengthOfLongestSubstring(s string) int {
    length := 0
    left := 0
    hash := make(map[byte]int)

    for right := 0; right < len(s); right++ {
        c := s[right]
        if index, ok := hash[c]; ok && index >= left {
            left = index+1
        }
        hash[c] = right
        length = max(length, right-left+1)
    }

    return length
}