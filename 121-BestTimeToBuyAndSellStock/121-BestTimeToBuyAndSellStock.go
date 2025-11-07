// Last updated: 11/7/2025, 2:48:55 PM
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0    

    for i := 1; i < len(prices); i++ {
        if currProfit := prices[i] - minPrice; currProfit > maxProfit {
            maxProfit = currProfit
        }
        if prices[i] < minPrice {
            minPrice = prices[i]
        }
    }

    return maxProfit
}