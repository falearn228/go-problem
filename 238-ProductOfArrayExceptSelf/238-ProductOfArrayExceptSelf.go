// Last updated: 11/7/2025, 2:48:11 PM
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1

	for i := range answer {
		if i > 0 {
			answer[i] = answer[i-1] * nums[i-1] // накапливаем префиксное произведение
		}
	}
    
    i := len(nums)-2 // т.к. у последнего элемента нет суфиксов
    sfx := nums[len(nums)-1] // поэтому суффикс накапливать начинаем с посл. элемента

    for i >= 0 {
        answer[i] *= sfx
        sfx *= nums[i]
        i -= 1
    }
    return answer
}