func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, len(nums))
	answer[0] = 1

	for i := range answer {
		if i > 0 {
			answer[i] = answer[i-1] * nums[i-1]
		}
	}

	//sufx[i] = sufx[i+1] * nums[i+1]
	i := n - 2
	sufx := nums[n-1]
	fmt.Println(answer)
	for i >= 0 {
		answer[i] *= sufx
		sufx = sufx * nums[i]
		i--
	}
	return answer
}