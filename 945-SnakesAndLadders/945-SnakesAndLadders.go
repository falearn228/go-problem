// Last updated: 12/2/2025, 4:35:31 PM
type Node struct {
    pos   int
    steps int
}

func numToSquare(num, n int) (int, int) {
    num-- //0-index
    rowFromBottom := num / n
    col := num % n

    row := n - rowFromBottom - 1

    if rowFromBottom % 2 == 1 {
        col = n - col - 1
    }

    return row, col
}

func snakesAndLadders(board [][]int) int {
    n := len(board)
    visited := make([]bool , n*n+1)
    visited[1] = true
    
    queue := []Node{{pos: 1, steps: 0}}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        if cur.pos == n*n {
            return cur.steps
        }

        for i := cur.pos+1; i <= min(cur.pos+6, n*n); i++ {
            next := i
            row, col := numToSquare(next, n)
            if board[row][col] != -1 {
                next = board[row][col]
            }

            if !visited[next] {
                visited[next] = true
                queue = append(queue, Node{steps: cur.steps+1, pos: next})
            }
        }
    }

    return -1
}