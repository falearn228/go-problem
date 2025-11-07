// Last updated: 11/7/2025, 2:47:10 PM
func orangesRotting(grid [][]int) int {
    rows := len(grid)
    if rows == 0 {
        return 0
    }
    queue := make([][2]int, 0)
    freshOranges := 0

    for i, arr := range grid {
        for j := range arr {
            if grid[i][j] == 2 {
                queue = append(queue, [2]int{i, j})
            } else if grid[i][j] == 1 {
                freshOranges++
            }
        }
    }

    if freshOranges == 0 {
        return 0
    }

    moves := [4][2]int {
        [2]int{1, 0},
        [2]int{-1, 0},
        [2]int{0, 1},
        [2]int{0, -1},
    }

    minutes := 0

    endRow := len(grid)
    endCol := len(grid[0])
    for len(queue) > 0 {
        lenQ := len(queue)

        rottenInThisMin := false

        for i := 0; i < lenQ; i++ {
            cell := queue[0]
            queue = queue[1:]

            for _, move := range moves {
                rowCell, colCell := cell[0] + move[0], cell[1] + move[1]

                if rowCell < 0 || colCell < 0 || rowCell >= endRow || colCell >= endCol {
                    continue
                }

                if grid[rowCell][colCell] == 1 {
                    rottenInThisMin = true
                    freshOranges--
                    grid[rowCell][colCell] = 2
                    queue = append(queue, [2]int{rowCell, colCell})
                }

            }
        }
        if rottenInThisMin {
            minutes++
        }
    }

    if freshOranges > 0 {
        return -1
    }

    return minutes
}