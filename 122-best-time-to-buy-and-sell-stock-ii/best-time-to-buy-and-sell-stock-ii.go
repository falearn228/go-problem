func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    maxProfit := 0
    
    for i := 1; i < len(prices); i++ {
        maxProfit += max(0, prices[i] - prices[i-1])
    }

    return maxProfit
}