// Last updated: 11/7/2025, 2:48:08 PM
func hIndex(citations []int) int {
    n := len(citations)
    arr := make([]int, n+1)

    for i := range citations {
        if citations[i] > n {
            arr[n]++ 
            continue
        }
        arr[citations[i]]++
    }

    totalzalups := 0
    for i := n; i >= 0; i-- {
        totalzalups += arr[i]
        if totalzalups >= i {
            return i
        }
    }
    return 0
}

// 0 1 3 5
// 2 <= 3