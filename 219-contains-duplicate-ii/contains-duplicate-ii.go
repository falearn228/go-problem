func containsNearbyDuplicate(nums []int, k int) bool {
    zalupa := make(map[int]int)

    for i, v := range nums {
        if j, ok := zalupa[v]; ok && abs(i, j, k) {
            return true
        }
        zalupa[v] = i
    }
    return false
}

func abs(i, j, k int) bool {
    if i >= j {
        if i-j <= k {
            return true
        }
    } else {
        if j-i <= k {
            return true 
        }
    }
    return false
}