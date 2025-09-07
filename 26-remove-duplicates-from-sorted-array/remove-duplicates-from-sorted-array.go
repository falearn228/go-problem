func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 1 {
        return 1
    }

    left, right := 1, 1

    for left < n {
        if nums[left] != nums[left-1] {
            nums[right] = nums[left]
            right++
        } 
        left++
    }

    return right
}