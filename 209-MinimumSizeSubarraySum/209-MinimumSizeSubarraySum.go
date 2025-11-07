// Last updated: 11/7/2025, 2:48:18 PM
func minSubArrayLen(target int, nums []int) int {
    left := 0
    sum := 0
    ans := 1_000_000_000
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum >= target {
            ans = min(ans, right - left+1)
            sum -= nums[left]
            left++
        }
    }
    if ans == 1_000_000_000 {
        return 0
    }
    return ans
}