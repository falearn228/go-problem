// Last updated: 11/7/2025, 2:49:20 PM
func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums)-1

    for left <= right {
        middle := (left + right) / 2

        value := nums[middle]
        if value == target {
            return middle
        } else if value > target {
            right = middle - 1
        } else {
            left = middle + 1
        }
    }

    return left
}