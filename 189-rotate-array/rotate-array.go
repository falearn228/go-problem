func rotate(nums []int, k int)  {
    n := len(nums)
    if n == 1 || k == 0 {
        return
    }

    k %= n 
    count, start := 0, 0

    for ; count < n; {
        nextIdx := start
        temp := nums[start]
        for {
            nextIdx = (nextIdx+k) % n
            temp, nums[nextIdx] = nums[nextIdx], temp
            count++
            if start == nextIdx {
                break
            }
        }
        start++
    }
}