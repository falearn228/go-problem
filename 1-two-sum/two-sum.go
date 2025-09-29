func twoSum(nums []int, k int) []int {
    zalupa := make(map[int]int, len(nums))

    for i := range nums {
        value := nums[i]
        if index, ok := zalupa[value]; ok && i != index {
            return []int{index, i}
        }
        zalupa[k-value]=i
    }

    return []int{}
}