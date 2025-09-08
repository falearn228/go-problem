func rotate(nums []int, k int)  {
    n := len(nums)
    if n == 1 || k == 0 {
        return
    }

    arr := make([]int, n)

    for i := range arr {
        newpos := (k + i) % n
        arr[newpos] = nums[i]
    }
    
    copy(nums, arr)
}