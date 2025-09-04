func maxProfit(prices []int, fee int) int {
    var hold, free int
    hold = -prices[0]

    for i := range prices {
        if i == 0 {
            continue
        }
        prevHold := hold
        hold = max(prevHold, free - prices[i])
        free = max(free, prevHold + prices[i] - fee)
    }

    return free
}