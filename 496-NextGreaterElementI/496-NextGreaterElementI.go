// Last updated: 11/7/2025, 2:47:47 PM
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    output := make([]int, len(nums1))

    hash := make(map[int]int, len(nums2) / 2)
    stack := make([]int, 0)

    for i := range nums2 {
        for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
            popped := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            hash[popped] = nums2[i]
        }
        stack = append(stack, nums2[i])
    }

    for i, v := range nums1 {
        val, ok := hash[v]
        if !ok {
            output[i] = -1 
        } else {
            output[i] = val
        }
    }   
    return output
}