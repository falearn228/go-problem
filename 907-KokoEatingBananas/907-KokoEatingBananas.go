// Last updated: 11/7/2025, 2:47:14 PM
func minEatingSpeed(piles []int, h int) int {
    var hoursNeeded func(mid int) int
    hoursNeeded = func(mid int) int {
        sum := 0
        
        for _, v := range piles {
            sum += (v + mid - 1) / mid
        }
        return sum
    }

    left := 1
    right := findMaxVal(piles)

    for left <= right {
        mid := left + (right - left)/2

        if hoursNeeded(mid) <= h {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}

func findMaxVal(piles []int) int {
    res := 1
    for _, v := range piles {
        if v > res {
            res = v
        }
    }
    return res
}