func majorityElement(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    
    candidate := 0
    count := 0

    for i := range nums {
        if count == 0 {
            candidate = nums[i]
        }
        if nums[i] == candidate {
            count++
        } else if nums[i] != candidate {
            count--
        }
    }

    return candidate
}