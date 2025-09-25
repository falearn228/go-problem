func rotate(matrix [][]int)  {
    n := len(matrix)
    // step 1. a -> a^T
    for i := 0; i < n; i++ {
        for j := i ; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    // step 2. reverse column
    for row := 0; row < n; row++ {
        for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
            matrix[row][i], matrix[row][j] = matrix[row][j], matrix[row][i]
        }
    }
}