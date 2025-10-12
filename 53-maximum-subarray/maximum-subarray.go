func maxSubArray(nums []int) int {
    maxCurr := nums[0]
    maxGlobal := nums[0]
    
    for i := 1; i < len(nums); i++ {
        num := nums[i]
        maxCurr = max(num, maxCurr + num)
        if maxCurr > maxGlobal {
            maxGlobal = maxCurr
        }
    }
    return maxGlobal
}