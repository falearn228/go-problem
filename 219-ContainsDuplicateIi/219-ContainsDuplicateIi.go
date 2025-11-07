// Last updated: 11/7/2025, 2:48:15 PM
func containsNearbyDuplicate(nums []int, k int) bool {
    zalupa := make(map[int]int)

    for i, v := range nums {
        if j, ok := zalupa[v]; ok && i-j <= k {
            return true
        }
        zalupa[v] = i
    }
    return false
}
