// Last updated: 11/7/2025, 2:48:52 PM
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxCount := 1
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			count++
		} else {
			maxCount = max(maxCount, count)
			count = 1
		}
	}
	return max(maxCount, count)
}