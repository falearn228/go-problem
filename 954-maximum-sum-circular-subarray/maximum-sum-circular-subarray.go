func maxSubarraySumCircular(nums []int) int {
    minCurr := nums[0]
    minGlobal := nums[0]
    maxKad := nums[0]
    maxKadCurr := nums[0]
    sum := nums[0]
    n := len(nums)
    for i := 1; i < n; i++ {
        num := nums[i]
        maxKadCurr = max(maxKadCurr + num, num)
        maxKad = max(maxKad, maxKadCurr) 
        minCurr = min(minCurr + num, num)
        minGlobal = min(minGlobal, minCurr) 
        sum += num
    }

    if sum == minGlobal {
        return maxKad
    }

    maxGlobal := sum - minGlobal


    return max(maxGlobal, maxKad)
}

// 1 2 3 1 2 3