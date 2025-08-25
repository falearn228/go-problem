func findMinArrowShots(points [][]int) int {
    if len(points) == 1 {
        return 1
    }

    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })

    lastEnd := points[0][1]
    answer := 1
    for i := 1; i < len(points); i++ {
        start := points[i][0]
        if lastEnd < start {
            lastEnd = points[i][1]
            answer++
        }
    }
    return answer
}