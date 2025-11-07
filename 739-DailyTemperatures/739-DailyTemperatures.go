// Last updated: 11/7/2025, 2:47:23 PM
func dailyTemperatures(temperatures []int) []int {
    if len(temperatures) == 1 {
        return []int{0}
    }
    answer := make([]int, len(temperatures))
    stack := make([]int, 0, len(temperatures))
    // stack[0] < stack[last] , return len, 

    for i := 0; i < len(temperatures); i++ {
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            lastInd := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            answer[lastInd] = i - lastInd
        }
        stack = append(stack, i)
    }
    return  answer
}