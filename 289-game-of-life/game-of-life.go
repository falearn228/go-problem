func gameOfLife(board [][]int)  {
    // cell was alive but next zdoxla -> -1
    // cell was zdoxla but next alive -> 2

    neighs := [8][]int{
        []int{-1, -1},
        []int{-1, 0},
        []int{-1, 1},
        []int{0, -1},
        []int{0, 1},
        []int{1, -1},
        []int{1, 0},
        []int{1, 1},
    }

    rows := len(board)
    cols := len(board[0])
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            cell := board[i][j]
            alive := 0
            for _, v := range neighs {
                ii, jj := i + v[0], j + v[1]
                if ii >= 0 && ii < rows && jj >= 0 && jj < cols {
                    neighbor := board[ii][jj]
                    if neighbor == 1 || neighbor == -1 {
                        alive++
                    }
                }
            }
            if cell == 1 {
                if alive < 2 || alive > 3 {
                    board[i][j] = -1
                } 
            } else if cell == 0 {
                if alive == 3 {
                    board[i][j] = 2
                }
            }
        }
    }

     for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 2 {
                board[i][j] = 1
            } else if board[i][j] == -1 {
                board[i][j] = 0
            }
        }
     }
}