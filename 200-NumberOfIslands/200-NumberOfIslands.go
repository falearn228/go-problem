// Last updated: 11/12/2025, 5:35:32 PM
type Object struct {
    row int
    col int
}

func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    n := len(grid)
    m := len(grid[0])
    
    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= n || c < 0 || c >= m || grid[r][c] != '1' {
            return
        }

        if grid[r][c] == '1' {
            grid[r][c] = '0'
        }

        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)
    }

    cnt := 0
    for i := range n {
        for j := range m {
            if grid[i][j] == '1' {
                cnt++
                dfs(i, j)
            }
        }
    }

    return cnt
}