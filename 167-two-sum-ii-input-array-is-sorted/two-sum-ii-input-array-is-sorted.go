func twoSum(numbers []int, target int) []int {
    ans := []int{1, 2}
    for i, v := range numbers {
        ind, ok := slices.BinarySearch(numbers, target-v)
        if ok {
            if i == ind {
                ind++
            }
            ans = []int{i+1, ind+1}
            break
        }
    }
    return ans
}