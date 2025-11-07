// Last updated: 11/7/2025, 2:49:09 PM
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0, len(intervals)+1)

	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}

	// 1 6  5 10
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	ans = append(ans, newInterval)

	for i < len(intervals) {
		ans = append(ans, intervals[i])
		i++
	}

	return ans
}