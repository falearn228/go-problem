// Last updated: 11/7/2025, 2:47:54 PM
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    var dfs func(startNode, endNode string, value float64) float64
    var visited map[string]bool
    neighbors := make(map[string]map[string]float64)
    for i := range equations {
        parrent := equations[i][0]
        child := equations[i][1]
         if neighbors[parrent] == nil {
            neighbors[parrent] = make(map[string]float64)
        }
        if neighbors[child] == nil {
            neighbors[child] = make(map[string]float64)
        }
        neighbors[parrent][child] = float64(values[i])
        neighbors[child][parrent] = 1.0 / values[i]
    }

    dfs = func(startNode, endNode string, value float64) float64 {
        if startNode == endNode {
            return value
        }

        visited[startNode] = true
        childs := neighbors[startNode]

        for chChild, chValue := range childs {
            if visited[chChild] {
                continue
            }
            result := dfs(chChild, endNode, value * chValue)
            if result != -1.0 {
                return result
            }
        }
        return -1.0
    }

    ans := make([]float64, len(queries))
    for i := range queries {
        curr := queries[i]
        if _, ok := neighbors[curr[0]]; !ok {
            ans[i] = -1.0
            continue
        }
        if _, ok := neighbors[curr[1]]; !ok {
            ans[i] = -1.0
            continue
        }
        visited = make(map[string]bool)
        res := dfs(curr[0], curr[1], 1.0)
        visited = nil
        ans[i] = res
    }

    return ans
}