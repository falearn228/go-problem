func productExceptSelf(nums []int) []int {
    prefixes := make([]int, len(nums))
    sufixes := make([]int, len(nums))
    n := len(nums)
    prefixes[0] = 1
    sufixes[n-1] = 1

    i, j := 1, n-2
    for i < n && j >= 0 {
        prefixes[i] = prefixes[i-1] * nums[i-1]
        sufixes[j] = sufixes[j+1] * nums[j+1]
        i++
        j--
    }

    answer := make([]int, len(nums))
    for i := range nums {
        answer[i] = prefixes[i] * sufixes[i]
    }
    return answer
}