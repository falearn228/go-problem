func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	ans := make([][]int, 0, len(intervals))

	left := 0
	right := len(intervals) - 1
	for left <= right {
		mid := (left + right) / 2

		if intervals[mid][0] > newInterval[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	for i := range left {
		ans = append(ans, intervals[i])
	}
	// ans = append(ans, newInterval)

	//вставка newInterval
	var last []int
	if len(ans) > 0 {
		last = ans[len(ans)-1]
		if newInterval[0] <= last[1] {
			if newInterval[1] > last[1] {
				ans[len(ans)-1][1] = newInterval[1]
			}
		} else {
			ans = append(ans, newInterval)
		}
	} else {
		ans = append(ans, newInterval)
	}
	
	// merge
	for i := left; i < len(intervals); i++ {
		last = ans[len(ans)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				ans[len(ans)-1][1] = curr[1]
			}
		} else {
			ans = append(ans, curr)
		}
	}

	return ans
}