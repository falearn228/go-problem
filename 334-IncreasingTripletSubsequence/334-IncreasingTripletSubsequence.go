// Last updated: 11/7/2025, 2:48:03 PM
func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
        return false
    }

    first := math.MaxInt32
    second := math.MaxInt32

    for i := 0; i < len(nums); i++ {
        if nums[i] > second {
            return true
        } else if nums[i] <= first {
            first = nums[i]
        } else if nums[i] <= second {
            second = nums[i]
        } 
    }
    return false
}