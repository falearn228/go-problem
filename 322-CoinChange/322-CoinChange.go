// Last updated: 12/24/2025, 12:17:57 PM
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1) // мин. кол-во монет для достижения суммы
    dp[0] = 0
    for i := 1; i < len(dp); i++ {
        dp[i] = math.MaxInt
    }

    for x := range amount+1 {
        for _, coin := range coins {
            // coin+dp[x-coin] = dp[x]
            if x-coin >= 0 && dp[x-coin] != math.MaxInt {
                dp[x] = min(dp[x], dp[x-coin]+1)
            }
        }
    }

    if dp[amount] == math.MaxInt {
        return -1
    }
    return dp[amount]
}

// 0 .. 11
//  1 2 11
//  x=3
//  3-1 >=0 , dp[2] = 1
//  dp[3] = min(dp[3], dp[2]+1)