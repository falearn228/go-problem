func equalPairs(grid [][]int) int {
    count := 0
    res := make([]int, len(grid))

    for i := 0; i < len(grid); i++ {
        for col := 0; col < len(grid); col++ {
            for j := range res {
                res[j] =  grid[j][col]
            }
            if slices.Equal(grid[i], res) {
                count += 1
            }
        }
    }
    
    return count
}