// Last updated: 11/7/2025, 2:49:16 PM
func permute(nums []int) [][]int {
	output := make([][]int, 0)
	used := make([]bool, len(nums))
	currPermut := make([]int, 0)

	sad := backtrack(&currPermut, &nums, used, &output)
	return *sad
}

func backtrack(currPermut, nums *[]int, used []bool, output *[][]int) *[][]int {
	if len(*currPermut) == len(*nums) {
		copyPer := make([]int, len(*currPermut))
		copy(copyPer, *currPermut)
		*output = append(*output, copyPer)
		return output
	}

	for i := range len(*nums) {
		if used[i] == false {
			*currPermut = append(*currPermut, (*nums)[i])
			used[i] = true

			output = backtrack(currPermut, nums, used, output)

			*currPermut = (*currPermut)[:len(*currPermut)-1]
			used[i] = false
		}

	}

	return output
}