// Last updated: 11/7/2025, 2:48:07 PM
func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	n := len(nums)

	write := 0
	for _, v := range nums {
		if v != 0 {
			nums[write] = v
			write++
		}
	}
	for write < n {
		nums[write] = 0
		write++
	}
}