// Last updated: 11/7/2025, 2:46:04 PM
func equalPairs(grid [][]int) int {
	count := 0
	hashRow := make(map[string]int, len(grid))

	var sb strings.Builder
	sb.Grow(len(grid) * 7)
	for indarr := range grid {
		for _, val := range grid[indarr] {
			sb.WriteByte('#')
			sb.WriteString(strconv.Itoa(val))
		}
		hashRow[sb.String()]++
		sb.Reset()
	}

	for row := range grid {
		for col := range grid[row] {
			sb.WriteByte('#')
			sb.WriteString(strconv.Itoa(grid[col][row]))
		}
		if _, ok := hashRow[sb.String()]; ok {
			count += hashRow[sb.String()]
		}
		sb.Reset()
	}

	return count
}