// Last updated: 11/7/2025, 2:46:17 PM
func nearestExit(maze [][]byte, entrance []int) int {
    queue := make([][2]int, 1)
    queue[0][0], queue[0][1] = entrance[0], entrance[1]

    maze[entrance[0]][entrance[1]] = '+'

    steps := 0
    moves := [4][2]int{
        [2]int{1, 0},
        [2]int{0, 1},
        [2]int{-1, 0},
        [2]int{0, -1},
    }

    endRow := len(maze)
    endCol := len(maze[0])

    for len(queue) > 0 {
        steps++

        lenQ := len(queue)
        for i := 0; i < lenQ; i++ {
            pos := queue[0]
            queue = queue[1:]

            for _, move := range moves {
                row := pos[0] + move[0]
                col := pos[1] + move[1]

                if row < 0 || row >= endRow || col < 0 || col >= endCol {
                    continue
                }

                if maze[row][col] == '+' {
                    continue
                }

                if row == 0 || row == endRow-1 || col == 0 || col == endCol-1 {
                    return steps
                }

                maze[row][col] = '+'
                queue = append(queue, [2]int{row, col})
            } 
        }
    }
    
    return -1
}