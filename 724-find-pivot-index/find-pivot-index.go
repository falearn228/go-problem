func pivotIndex(nums []int) int {
    pivotIndex := -1

    // 0 1 8 11 17 22 28
    // ps[n-1] - ps[4] == ps[3] -> nums[3]

    prefSum := make([]int, len(nums)+1)
    n := len(prefSum)

    for i := 1; i < n; i++ {
        prefSum[i] = prefSum[i-1] + nums[i-1]
    }

    for i := n-1; i > 0; i-- {
        if prefSum[n-1] - prefSum[i] == prefSum[i-1] {
            pivotIndex = i-1
        } 
    }
    return pivotIndex
}