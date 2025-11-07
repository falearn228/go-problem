// Last updated: 11/7/2025, 2:47:53 PM
func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 1 {
        return 0
    }

    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    lastEnd := intervals[0][1]
    ans := 0
    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]
        if start >= lastEnd {
            lastEnd = intervals[i][1]
        } else {
            ans++
        }
    }
    return ans
}