// arr1[end] >= arr2[start] 
// 1,4   4,7
func merge(intervals [][]int) [][]int {
    slices.SortFunc(intervals, func(i, j []int) int {
		return int(i[0] - j[0])
	})

    ans := make([][]int, 0, len(intervals))
    ans = append(ans, intervals[0])
    for i := 1; i < len(intervals); i++ {
        last := ans[len(ans)-1]
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