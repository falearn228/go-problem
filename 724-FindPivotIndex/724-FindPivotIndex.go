// Last updated: 11/7/2025, 2:47:26 PM
func pivotIndex(nums []int) int {
    var leftSum, rightSum int

    for _, v := range nums {
        rightSum += v
    } 
    for i, v := range nums {
        rightSum -= v
        if leftSum == rightSum {
            return i
        }
        leftSum += v
    }
    return -1
}