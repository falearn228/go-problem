// Last updated: 11/7/2025, 2:49:03 PM
func setZeroes(matrix [][]int)  {
    rows := len(matrix)
    cols := len(matrix[0])
    isFirstColZero := false

    for i := 0; i < rows; i++ {
        if matrix[i][0] == 0 {
            isFirstColZero = true
        }

        for j := 1; j < cols; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i := rows-1; i >= 0; i-- {
        for j  := cols-1; j >= 1; j-- {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    if isFirstColZero {
        for i := range rows {
            matrix[i][0] = 0
        }
    }

}

func init (){
    debug.SetMemoryLimit(7)
}