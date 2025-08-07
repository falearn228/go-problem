func uniqueOccurrences(arr []int) bool {
    hash := make(map[int]int)
    set := make(map[int]struct{})

    for _, v := range arr {
        hash[v]++
    }
    for _, value := range hash {
        if _, ok := set[value]; ok {
            return false
        }
        set[value] = struct{}{}
    }
    return true
}