// Last updated: 11/7/2025, 2:49:12 PM
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    rows := len(matrix)
    cols := len(matrix[0])
    ans := make([]int, 0, rows*cols)

    top, bottom := 0, rows-1
    left, right := 0, cols-1

    for top <= bottom && left <= right {
        for j := left; j <= right; j++ {
            ans = append(ans, matrix[top][j])
        }
        top++

        if top > bottom {
            break
        }

        for i := top; i <= bottom; i++ {
            ans = append(ans, matrix[i][right])
        }
        right--

        if right < left {
            break
        }

        for j := right; j >= left; j-- {
            ans = append(ans, matrix[bottom][j])
        }
        bottom--

        for i := bottom; i >= top; i-- {
            ans = append(ans, matrix[i][left])
        }
        left++
    }

    return ans
}