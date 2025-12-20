func isInterleave(s1 string, s2 string, s3 string) bool {
    n := len(s1)
    m := len(s2)
    if n < m {
        return isInterleave(s2,s1,s3)
    }
    if n + m != len(s3) {
        return false
    }

    dp := make([]bool, m+1)

    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            if i == 0 && j == 0 {
                dp[i] = true
                continue
            }
            if i == 0 {
                // s2 берем, мы пришли слева, то есть по сути это новое значение
                dp[j] = dp[j-1] && (s2[j-1] == s3[j-1]) 
            } else if j == 0 {
                // s1, мы пришли сверху, тк у нас одномерный массив, то ничего не изменится
                // ну а если был бы двумерный, то стало типо dp[i-1][j]
                dp[j] = dp[j] && (s1[i-1] == s3[i-1])
            } else {
                dp[j] = (dp[j-1] && s2[j-1] == s3[i+j-1]) || (dp[j] && s1[i-1] == s3[i+j-1])
            }
        }
    }

    return dp[m]
}