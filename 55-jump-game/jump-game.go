func canJump(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return true
    }

    var currjump int
    for ind := 0; ind <= currjump; ind++ {
        for start := ind; start < nums[ind] + ind; start++ {
            currjump = max(nums[start] + start, currjump)
            if currjump >= n-1 {
                return true
            }
        }
    }

    return false
}