func reverseWords(s string) string {
    fields := strings.Fields(s)

    startPtr, endPtr := 0, len(fields) - 1

    for startPtr < endPtr {
        fields[startPtr], fields[endPtr] = fields[endPtr], fields[startPtr]
        
        startPtr++
        endPtr--
    }

    return strings.Join(fields, " ")
}