func trap(height []int) int {
    maxLefts := make([]int, len(height))
    maxRights := make([]int, len(height))

    for i := range height {
        for j := 0; j < i; j++ {
            maxLefts[i] = max(maxLefts[i], height[j]) 
        }
    }

    for i := range height {
        for j := len(height)-1; j > i; j-- {
            maxRights[i] = max(maxRights[i], height[j]) 
        }
    }

    var water int
    for i := 0; i < len(height); i++ {
        waterHeight := min(maxRights[i], maxLefts[i])
        diff := waterHeight - height[i]
        if diff > 0 {
            water += diff
        }
    }

    return water
}