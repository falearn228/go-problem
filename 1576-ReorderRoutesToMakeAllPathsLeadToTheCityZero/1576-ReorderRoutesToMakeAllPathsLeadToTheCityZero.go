// Last updated: 11/7/2025, 2:46:36 PM
func minReorder(n int, connections [][]int) int {
    stack := make([]int, 1)
    visited := make(map[int]bool)
    neighbors := make(map[int][]int)
    cons := make(map[int][]int)
    for _, v := range connections {
        route := v[0]
        neigh := v[1]
        neighbors[route] = append(neighbors[route], neigh)
        neighbors[neigh] = append(neighbors[neigh], route)
        cons[route] = append(cons[route], neigh)
    }

    var res int
    for len(stack) > 0 {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if visited[top] {
            continue
        }
        visited[top] = true

        neighs := neighbors[top]
        for i := range neighs {
            parrent := top
            child := neighs[i]
	        if visited[child] {
				continue
			}
            have := slices.Contains(cons[parrent], child)
            
            if have {
                res++
            }

            stack = append(stack, child)
        }
    }
    return res
}