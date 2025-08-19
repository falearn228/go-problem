func successfulPairs(spells []int, potions []int, success int64) []int {
    n := len(spells)
    m := len(potions)
    pairs := make([]int, n)

    // mergeSort(potions) NlogN
    slices.Sort(potions)
    // find index -> 1 2 3 4 5 6 7 8 9 10
    for i, v := range spells {
        left := 0
        right := m - 1
        countPairs := 0 
        for left <= right {
            mid := left + (right - left) / 2

            power := int64(potions[mid]) * int64(v)
            if power >= success {
                // pairs[i] = mid
                countPairs = m - mid
                right = mid - 1
            } else {
                left = mid + 1
            } 
        }
        pairs[i] = countPairs
    }

    return pairs
}