// Last updated: 11/7/2025, 2:49:19 PM
func isValidSudoku(board [][]byte) bool {
    rows := make([]map[byte]bool, 9)
    cols :=  make([]map[byte]bool, 9)
    squares := make([]map[byte]bool, 9)
    
    for i := 0; i < len(rows); i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        squares[i] = make(map[byte]bool)
    }

    for i := range board {
        for j := range board[0] {
            num := board[i][j]
            if num == '.' || num == ',' {
                continue
            } 

            if rows[i][num] == true {
                return false
            } 
            rows[i][num] = true

            if cols[j][num] == true {
                return false
            } 
            cols[j][num] = true

            box := (i/3)*3 + j/3 
            if squares[box][num] == true {
                return false
            } 
            squares[box][num] = true
        }

    }

    return true
}