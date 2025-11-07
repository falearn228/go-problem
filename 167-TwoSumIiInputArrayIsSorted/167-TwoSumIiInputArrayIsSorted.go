// Last updated: 11/7/2025, 2:48:39 PM
func twoSum(numbers []int, target int) []int {
    ans := []int{1, 2}
    i, j := 0, len(numbers)-1
    for i < j {
        if numbers[j] + numbers[i] > target {
            j--
        } else if numbers[j] + numbers[i] < target {
            i++
        } else {
            return []int{i+1, j+1}
        }
    }
    return ans
}