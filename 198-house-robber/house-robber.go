func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    robs := make([]int, len(nums)+1)
    robs[1] = nums[0]

    for i := 1; i < len(nums); i++ {
        robs[i+1] = max(robs[i], robs[i-1]+nums[i])
    } 
    return robs[len(robs)-1]
}