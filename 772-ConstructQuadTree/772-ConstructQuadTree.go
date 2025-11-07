// Last updated: 11/7/2025, 2:47:20 PM
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    var dfs func(x, y, size int) *Node  

    dfs = func(x, y, size int) *Node {
        isEqual := true
        firstVal := grid[x][y]
l:
        for i := x; i < x+size; i++ {
            for j := y; j < y+size; j++ {
                if grid[i][j] != firstVal {
                    isEqual = false
                    break l
                }
            }
        }

        if isEqual {
            return &Node{Val: firstVal == 1, IsLeaf: true}
        }

        newSize := size / 2
        node := &Node{
            IsLeaf: false,
            Val: true,
        }

        node.TopLeft  = dfs(x, y, newSize)
        node.TopRight = dfs(x, y+newSize, newSize)
        node.BottomLeft = dfs(x+newSize, y, newSize)
        node.BottomRight = dfs(x+newSize, y+newSize, newSize)

        return node
    }

    root := dfs(0, 0, len(grid))
    return root
}