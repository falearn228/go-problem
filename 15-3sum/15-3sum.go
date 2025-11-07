// Last updated: 11/7/2025, 2:49:31 PM
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    ans := make([][]int, 0)
    n := len(nums)

    for i := 0; i < n-2; i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        } 

        left := i+1
        right := n-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            switch {
            case sum == 0:
                ans = append(ans, []int{nums[i], nums[left], nums[right]})
                lv := nums[left]
                for left < right && nums[left] == lv {
                    left++
                } 

                rv := nums[right]
                for left < right && nums[right] == rv {
                    right--
                }
            case sum < 0:
                left++
            case sum > 0:
                right--
            }

        }
    }

    return ans
}

// -4 -1 -1 0 1 2