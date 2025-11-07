// Last updated: 11/7/2025, 2:49:22 PM
func strStr(haystack string, needle string) int {
    nlen := len(needle)
    hlen := len(haystack)
    if hlen < nlen {
        return -1
    }

    // O(n*m) - интуитивное решение
    // for i := 0; i <= hlen-nlen; i++ {
    //     j := 0
    //     for j < nlen && haystack[i+j] == needle[j] {
    //         j++
    //     }

    //     if j == nlen {
    //         return i
    //     }
    // }
    // return -1

    lps := make([]int, nlen)
    length := 0
    i := 1
    for i < nlen {
        if needle[i] == needle[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                lps[i] = 0
                i++
            }
        }
    }

    i = 0
    j := 0
    for i < hlen {
        if needle[j] == haystack[i] {
            i++
            j++
        }

        if j == nlen {
            return i-j
        } else if i < hlen && needle[j] != haystack[i] {
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    return -1
}