func jump(nums []int) int {
    len := len(nums)
    if len == 1 {
        return 0
    }

    currEnd := 0
    count := 0
    farthest := 0

    for i := 0 ; i < len-1; i++ {
        if i+nums[i] > farthest {
            farthest = i + nums[i]
        }
        

        if i == currEnd  || farthest >= len-1{
            count++
            currEnd = farthest

            if currEnd >= len-1 {
                break
            }
        }
    }
    return count
}