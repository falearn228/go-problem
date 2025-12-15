func uniquePathsWithObstacles(g [][]int) int {
    if len(g) == 0 {
        return 0
    }

    // кол-во путей в текущей клетке
    n := len(g)
    m := len(g[0])
    dp := make([][]int, n)
    for j := range dp {
        dp[j] = make([]int, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            dp[i][j] = 1
            if g[i][j] == 1 {
                dp[i][j] = 0
                continue
            }
            if i == 0 && j == 0 {
                continue
            }
            if i == 0 {
                dp[i][j] = dp[i][j-1]
            } else if j == 0 {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }
        }
    }

    return dp[n-1][m-1]
}