// Last updated: 11/7/2025, 2:46:39 PM
func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := -1_000_000_000
    result := make([]bool, len(candies))
    for i := range candies {
        if candies[i] > max {
            max = candies[i]
        }
    }

    for i := range candies {
        if extraCandies + candies[i] >= max {
            result[i] = true
        }
    }
    return result
}