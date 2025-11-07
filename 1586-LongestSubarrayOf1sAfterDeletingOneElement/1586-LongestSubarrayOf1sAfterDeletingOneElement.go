// Last updated: 11/7/2025, 2:46:34 PM
func longestSubarray(nums []int) int {
    k := 1
    zeroNum := 0
    left := 0 
    maxLeng := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            zeroNum++
        }
        for zeroNum > k {
            if nums[left] == 0 {
                zeroNum--
            }
            left++
        }

        currLeng := right - left
        if currLeng > maxLeng {
            maxLeng = currLeng
        }
    }
    return maxLeng
}