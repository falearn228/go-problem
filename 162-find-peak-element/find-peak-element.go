func findPeakElement(nums []int) int {
    n := len(nums)
    left := 0
    right := n - 1

    for left <= right {
        mid := left + (right - left) / 2

        leftVal := math.MinInt
        if mid-1 >= 0 {
            leftVal = nums[mid-1]
        }
        rightVal := math.MinInt
        if mid+1 < n {
            rightVal = nums[mid+1]
        }

        currVal := nums[mid]
        if currVal > leftVal && currVal > rightVal {
            return mid
        }

        if currVal < rightVal {
            left = mid+1
        } else if currVal > rightVal  {
            right = mid-1
        } else if currVal < leftVal  {
            left = mid+1
        } else if currVal < leftVal  {
            right = mid-1
        }
    }
    return left
}