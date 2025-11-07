// Last updated: 11/7/2025, 2:46:27 PM
func maxOperations(nums []int, k int) int {
    sort.Ints(nums) // o n log n
    opers := 0
    n := len(nums)
    i, j := 0, n-1

    for i < j {
        if nums[i] + nums[j] < k {
            i++
        } else if nums[i] + nums[j] > k {
            j--
        } else {
            opers++
            i++
            j--
        }
    }
    return opers
}