func canJump(nums []int) bool {
    n := len(nums)
    if n == 1 {
        return true
    }

    var maxReach int
    for i, jump := range nums {
        if maxReach >= n-1  {
            return true
        }
        if i == n-1 || i > maxReach {
            return false
        }
        maxReach = max(maxReach, i + jump)
    }


    return true
}