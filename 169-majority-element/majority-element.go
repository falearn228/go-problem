func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    sort.Ints(nums)
    return nums[n / 2]
}