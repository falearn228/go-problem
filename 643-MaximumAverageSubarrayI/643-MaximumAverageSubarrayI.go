// Last updated: 11/7/2025, 2:47:29 PM
func findMaxAverage(nums []int, k int) float64 {
    maxAvg := math.Inf(-1)
    maxSum := 0
    for i := range k {
        maxSum += nums[i]
    }
    currSum := maxSum
    for i := k; i < len(nums); i++ {
        currSum = currSum + nums[i] - nums[i - k]
        if maxSum < currSum {
            maxSum = currSum
        }
    }

    maxAvg = float64(maxSum) / float64(k)

    return maxAvg
}
