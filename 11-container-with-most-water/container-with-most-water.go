func maxArea(height []int) int {
    // Y * X -> min(y1, y2) * (x2-x1) = 49

    n := len(height)
    if n == 2 {
        return min(height[1], height[0])
    }

    i, j := 0, n-1

    maxSquare := 0
    for i < j {
        square := min(height[i], height[j]) * (j-i)
        maxSquare = max(maxSquare, square)
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }

    return maxSquare
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}