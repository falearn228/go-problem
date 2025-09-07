func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }

    var k int
    tempNums := make([]int, len(nums))
    
    i1, i2 := 0, 0
    for i1 < len(nums) {
        if i2 >= 0 && nums[i1] != val {
            tempNums[i2] = nums[i1]
            i2++
        }
        i1++
    }

    k = i2
    for i := 0; i <= i2 && i2 < len(nums); i++ {
        nums[i] = tempNums[i]
    }


    return k
}