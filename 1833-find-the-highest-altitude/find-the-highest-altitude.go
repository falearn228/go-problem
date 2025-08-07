func largestAltitude(gain []int) int {
    var currAlt, maxAlt int
    for _, v := range gain {
        currAlt += v
        if currAlt > maxAlt {
            maxAlt = currAlt
        }
    }
    return maxAlt 
}
