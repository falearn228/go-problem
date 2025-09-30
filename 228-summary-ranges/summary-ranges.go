func summaryRanges(nums []int) []string {
	res := make([]string, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		j := i
		for ; j < len(nums)-1 && nums[j+1] == nums[j]+1; j++ {
		}

		if j == i {
			start := strconv.Itoa(nums[j])
			res = append(res, start)
		} else {
			start := strconv.Itoa(nums[i])
			end := strconv.Itoa(nums[j])

			res = append(res, start+"->"+end)
			i = j
		}

	}

	return res
}