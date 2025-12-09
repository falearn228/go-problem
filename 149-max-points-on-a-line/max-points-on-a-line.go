func maxPoints(points [][]int) int {
    n := len(points)
    if n <= 2 {
        return n
    }
    res := 1
    for i := 0; i < len(points); i++ {
        pointsMap := make(map[float64]int)
        for j := i+1; j < len(points); j++ {
            var k float64
            x1, y1 := float64(points[i][0]), float64(points[i][1])
            x2, y2 := float64(points[j][0]), float64(points[j][1])
            dx := x2 - x1
            dy := y2 - y1
            if dx == 0 {
                k = math.Inf(1)
            } else {
                k = dy/dx
            }

            pointsMap[k]++
            for _, counts := range pointsMap {
                if counts+1 > res {
                    res = counts+1
                }
            }
        }
    }

    return res
}

// k*x = y
// tgA = y'= k