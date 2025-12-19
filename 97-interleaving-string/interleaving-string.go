func isInterleave(s1 string, s2 string, s3 string) bool {
    n := len(s1)
    m := len(s2)
    if n + m != len(s3) {
        return false
    }

    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, m+1)
    }
    // для пустых строк true
    dp[0][0] = true

    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            if i > 0 {
                dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1])
            }
            if j > 0 && !dp[i][j] {
                dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1])
            }
        }
    }

    return dp[n][m]
}