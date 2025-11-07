// Last updated: 11/7/2025, 2:48:49 PM
func singleNumber(nums []int) int {
    var uniq_el int

    for i := range nums {
        uniq_el ^= nums[i]
    }
    return uniq_el
}