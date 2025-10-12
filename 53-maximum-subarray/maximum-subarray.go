func maxSubArray(nums []int) int {
    maxCurr := 0
    maxGlobal := nums[0]
    for i := range nums {
        maxCurr += nums[i]
        if maxCurr > maxGlobal {
            maxGlobal = maxCurr
        }
        if maxCurr <= 0 {
            maxCurr = 0
            continue
        }
    }

    return maxGlobal
}