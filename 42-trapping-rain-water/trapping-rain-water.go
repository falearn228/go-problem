func trap(height []int) int {
    n := len(height)
    if n <= 2 {
        return 0
    }
    left := 0
    right := len(height)-1

    maxLeft, maxRight := 0, 0
    totalWater := 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] > maxLeft {
                maxLeft = height[left]
            } else {
                totalWater += maxLeft - height[left]
            }
            left++
        } else {
            if height[right] > maxRight {
                maxRight = height[right]
            } else {
                totalWater += maxRight - height[right]
            }
            right--
        }
    }
    return totalWater
}