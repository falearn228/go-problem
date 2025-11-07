// Last updated: 11/7/2025, 2:49:02 PM
func minWindow(s string, t string) string {
    subLen := len(t)
    strLen := len(s)

    if subLen > strLen {
        return ""
    }

    minLen := ""
    byte_c := make(map[byte]int)
    for i := range t {
        byte_c[t[i]]++
    }

    left := 0 
    curr_byte_c := make(map[byte]int)
    curr_len := 0
    need_len := len(byte_c) // len keys -> unique symbs

    for right := left; right < strLen; right++ {
        char := s[right]
        if _, ok := byte_c[char]; !ok {
            continue
        }
        
        curr_byte_c[char]++

        if curr_byte_c[char] == byte_c[char] {
            curr_len++
        }

        for curr_len == need_len {
            if minLen == "" || (right-left+1) < len(minLen) {
                minLen = s[left:right+1]
            }

            left_c := s[left]

            if cnt, ok := byte_c[left_c]; ok {
                if curr_byte_c[left_c] == cnt {
                    curr_len--
                }
                curr_byte_c[left_c]--
            }

            left++
        }

    }
    

    return minLen
}