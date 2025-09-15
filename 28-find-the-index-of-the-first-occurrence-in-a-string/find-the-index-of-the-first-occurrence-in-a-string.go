func strStr(haystack string, needle string) int {
    nlen := len(needle)
    hlen := len(haystack)
    if hlen < nlen {
        return -1
    }

    for i := 0; i <= hlen-nlen; i++ {
        j := 0
        for j < nlen && haystack[i+j] == needle[j] {
            j++
        }

        if j == nlen {
            return i
        }
    }


    return -1
}