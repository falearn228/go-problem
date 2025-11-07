// Last updated: 11/7/2025, 2:49:01 PM
func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 1
    }

    insPos := 2

    for i := 2; i < n; i++ {
        if nums[i] != nums[insPos-2] {
            nums[insPos] = nums[i]
            insPos++
        }
    }
    return insPos
}