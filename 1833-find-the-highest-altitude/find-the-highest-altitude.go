func largestAltitude(gain []int) int {
    prefSum := make([]int, len(gain)+1)
    maxAltitude := 0
    for i := 1; i < len(prefSum); i++ {
        prefSum[i] = prefSum[i-1] + gain[i-1]
        maxAltitude = max(prefSum[i], maxAltitude)
    }

    return maxAltitude
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}