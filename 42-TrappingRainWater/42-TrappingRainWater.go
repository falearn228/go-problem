// Last updated: 11/7/2025, 2:49:18 PM
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
            maxLeft = max(maxLeft, height[left])
            totalWater += maxLeft - height[left]
            left++
        } else {
            maxRight = max(maxRight, height[right])
            totalWater += maxRight - height[right]
            right--
        }
    }
    return totalWater
}