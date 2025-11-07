// Last updated: 11/7/2025, 2:48:14 PM
func calculate(s string) int {
	stack := []int{}
	operand := 0
	result := 0
	sign := 1

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			operand = operand*10 + int(ch-'0')
			continue
		}
		switch ch {
		case '(':
			stack = append(stack, result)
			stack = append(stack, sign)

			result = 0
			sign = 1
		case ')':
			result += sign * operand
			operand = 0
			prevSign := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prevResult := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result *= prevSign
			result += prevResult
		case '+':
			// Выполняем предыдущую операцию
			result += sign * operand
			// Устанавливаем знак для следующей операции
			sign = 1
			operand = 0
		case '-':
			result += sign * operand
			sign = -1
			operand = 0

		}
	}

	return result + sign * operand
}