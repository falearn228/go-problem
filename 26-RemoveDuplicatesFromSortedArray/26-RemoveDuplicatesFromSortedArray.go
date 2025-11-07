// Last updated: 11/7/2025, 2:49:24 PM
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 1
    }

    right := 1

    for left := 1; left < n; left++ {
        if nums[left] != nums[left-1] {
            nums[right] = nums[left]
            right++
        } 
    }

    return right
}