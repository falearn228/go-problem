// Last updated: 11/7/2025, 2:47:09 PM
func longestOnes(nums []int, k int) int {
    zeroCount := 0
    left, right := 0, 0
    maxLeng := 0
    for right < len(nums) && left < len(nums) {
        if nums[right] == 0 {
            zeroCount++
        }
        if zeroCount <= k {
            currLen := right - left + 1
            maxLeng = max(maxLeng, currLen)
        } else {
            for zeroCount > k && left < len(nums) {
                if nums[left] == 0 {
                    zeroCount--
                }
                left++
            }
        }
        right++
    }
    return maxLeng
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}