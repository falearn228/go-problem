// Last updated: 11/7/2025, 2:49:23 PM
func removeElement(nums []int, val int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }


    curr1, curr2 := 0, n-1

    for curr1 <= curr2 {
        if nums[curr1] == val && nums[curr2] != val {
            nums[curr1] = nums[curr2]
            curr2--
            curr1++
        } else if nums[curr1] == val && nums[curr2] == val {
            curr2--
        } else {
            curr1++
        }
    }


    return curr1
}