// Last updated: 11/7/2025, 2:48:44 PM
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	operators := map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}

	for _, v := range tokens {
		if _, ok := operators[v]; !ok {
			intval, _ := strconv.Atoi(string(v))
			stack = append(stack, intval)
		} else {
			var ans int
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch v {
			case "+":
				ans += num2 + num1
			case "-":
				ans += num2 - num1
			case "*":
				ans += num2 * num1
			case "/":
				ans += num2 / num1
			}
			stack = append(stack, ans)
		}
	}

	return stack[0]
}
