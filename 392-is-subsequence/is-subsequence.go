func isSubsequence(s string, t string) bool {
    sPtr, tPtr := 0, 0

    // a b c
    // a h b g d c
    for tPtr < len(t) {
        if sPtr < len(s) && s[sPtr] == t[tPtr] {
            sPtr++
        } 
        tPtr++
    }

    return sPtr == len(s)
}