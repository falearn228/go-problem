// Last updated: 11/7/2025, 2:46:23 PM
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
