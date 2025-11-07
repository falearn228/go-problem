// Last updated: 11/7/2025, 2:49:30 PM
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    digitsMap := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    buffer := make([]byte, 0)
    res := make([]string, 0)
    var dfs func(int)
    dfs = func(ind int) {
        if ind == len(digits) {
            res = append(res, string(buffer))
            return
        }

        digit := digits[ind]
        chars, ok := digitsMap[digit]
        if !ok {
            dfs(ind+1)
            return
        }

        for i := 0; i < len(chars); i++ {
            buffer = append(buffer, chars[i])
            dfs(ind+1)
            if len(buffer) > 0 {
                buffer = buffer[:len(buffer)-1]
            }
        }
    }

    dfs(0)
    return res
}