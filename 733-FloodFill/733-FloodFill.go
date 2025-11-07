// Last updated: 11/7/2025, 2:47:25 PM
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	currColor := image[sr][sc]
	if currColor == color {
		return image
	}
	queue := make([][2]int, 0)
	visited := make(map[[2]int]bool, 0)
	queue = append(queue, [2]int{sr, sc})

	rows, cols := len(image), len(image[0])
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if flag := visited[top]; flag {
			continue
		}
		visited[top] = true
		row, col := top[0], top[1]
		if row < 0 || row >= rows || col < 0 || col >= cols || image[row][col] != currColor {
			continue
		}
		image[row][col] = color

		neighbors := [][2]int{
			{row - 1, col}, // вверх
			{row + 1, col}, // вниз
			{row, col - 1}, // влево
			{row, col + 1}, // вправо
		}

		for _, neighbor := range neighbors {
			nr, nc := neighbor[0], neighbor[1]
			// Проверяем границы и что еще не посещали
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return image
}
