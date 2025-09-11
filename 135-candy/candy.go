func candy(ratings []int) int {
    n := len(ratings)
    if n == 1 {
        return 1
    }
    
    candies := make([]int, n)

    for i := range candies {
        candies[i] = 1
        if i > 0 && ratings[i-1] < ratings[i] {
            candies[i] += candies[i-1]
        }
    }

    totalCandies, rightCandies := 0, 0
    for i := n-1; i >= 0; i-- {
        if i < n-1 && ratings[i+1] < ratings[i] {
            rightCandies++
        } else {
            rightCandies = 1
        }

        totalCandies += max(candies[i], rightCandies)
    }
    return totalCandies
}