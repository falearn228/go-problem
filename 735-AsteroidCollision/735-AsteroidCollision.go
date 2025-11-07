// Last updated: 11/7/2025, 2:47:24 PM
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		alive := true

		for alive && asteroids[i] < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top < -asteroids[i] {
				stack = stack[:len(stack)-1]
				continue
			} else if top+asteroids[i] == 0 {
				stack = stack[:len(stack)-1]
			}
			alive = false
		}
		if alive {
			stack = append(stack, asteroids[i])
		}
	}

	return stack
}