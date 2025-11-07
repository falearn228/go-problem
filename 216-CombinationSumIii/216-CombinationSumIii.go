// Last updated: 11/7/2025, 2:48:16 PM
func combinationSum3(k int, n int) [][]int {
    res := make([][]int, 0)
    buffer := make([]int, 0, k)

    var dfs func(start, currSum int)
    dfs = func(start, currSum int) {
        if len(buffer) == k {
            if currSum == n {
                combo := make([]int, len(buffer))
                copy(combo, buffer)
                res = append(res, combo)
            }
            return
        }

        for i := start; i < 10; i++ {
            if i > n - currSum {
                break
            }
            buffer = append(buffer, i)
            dfs(i+1, currSum+i)
            buffer = buffer[:len(buffer)-1]
        }
    }
    dfs(1, 0)
    return res
}