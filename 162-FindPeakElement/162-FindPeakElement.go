// Last updated: 11/7/2025, 2:48:40 PM
func findPeakElement(nums []int) int {
    l, r := 0, len(nums)-1
    for l < r {
        mid := l + (r-l)/2
        if nums[mid] > nums[mid+1] {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return l
}